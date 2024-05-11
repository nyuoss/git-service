package handler

import (
	"encoding/json"
	"fmt"
	"git-service/pkg/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	urlpkg "net/url"

	"github.com/gorilla/mux"
)

type CommitHandler interface {
	GetCommitsBefore(http.ResponseWriter, *http.Request)
	GetCommitsAfter(http.ResponseWriter, *http.Request)
	GetCommitByName(http.ResponseWriter, *http.Request)
	CommitReleased(http.ResponseWriter, *http.Request)
	GetJobsByCommit(http.ResponseWriter, *http.Request)
}

var _ CommitHandler = &commitHandler{}

type commitHandler struct{}

func NewCommitHandler() CommitHandler {
	return &commitHandler{}
}

func AddhttpAuthRequestHeaders(req *http.Request, personalAccessToken string) {
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+personalAccessToken)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
}

func (h *commitHandler) GetCommitsBefore(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *commitHandler) GetCommitsAfter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]

	queryParams := r.URL.Query()

	commitID := queryParams.Get("commit")
	number := queryParams.Get("number")

	commit, err := getCommit(owner, repo, commitID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	commitDateTime := commit.Author.Date
	datetime, err := time.Parse(time.RFC3339, commitDateTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := resty.New()
	var branches []Branch

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches", owner, repo)
	resp, err := client.R().SetResult(&branches).Get(apiURL)

	if err != nil {
		http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
		return
	}

	commitsByBranch := make(map[string][]string)
	additions := 0
	for _, branch := range branches {
		commits, err := getBranchCommits(owner, repo, branch.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var filteredCommits []string
		for _, commit := range commits {
			commitTime, err := time.Parse(time.RFC3339, commit.Commit.Author.Date)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if commitTime.After(datetime) {
				filteredCommits = append(filteredCommits, commit.SHA)
				additions++
			}

			
			if number != "" {
				num, err := strconv.Atoi(number)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if num == additions {
					break
				}
			}
		}

		if len(filteredCommits) > 0 {
			commitsByBranch[branch.Name] = filteredCommits
		}
	}

	response := GetCommitsAfterResp{
		Commits: commitsByBranch,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func getBranchCommits(owner, repo, branchName string) ([]model.CommitData, error) {
	var commits []model.CommitData
	client := resty.New()

	resp, err := client.R().
		SetResult(&commits).
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s", owner, repo, branchName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status code: %d", resp.StatusCode())
	}

	return commits, nil
}

func (h *commitHandler) GetCommitByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get request data from query params
	request, errMessage := GetCommitByNameRequest(r)
	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	branchExists, err := checkIfBranchExists(request.Owner, request.Repository, request.Branch, request.PersonalAccessToken)
	if err != nil {
		http.Error(w, "Error checking if branch exists: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !branchExists {
		http.Error(w, "Branch does not exist in the repository", http.StatusBadRequest)
		return
	}

	baseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s&per_page=100&page=", request.Owner, request.Repository, request.Branch)
	method := "GET"

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		http.Error(w, "Error generating new request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	AddhttpAuthRequestHeaders(req, request.PersonalAccessToken)

	resp := []model.CommitData{}
	var commits []model.CommitData
	client := &http.Client{}

	for page_number := 1; ; page_number++ {
		commits, errMessage = getCommitsByPageNumber(baseUrl, page_number, req, client)
		if errMessage != "" {
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}

		if len(commits) == 0 {
			break
		}

		for _, c := range commits {
			if strings.Contains(strings.ToLower(c.Commit.Message), strings.ToLower(request.CommitMessage)) {
				resp = append(resp, c)
			}
		}
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func (h *commitHandler) CommitReleased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get request data from query params
	request, errMessage := GetCommitReleasedRequest(r)
	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	branchExists, err := checkIfBranchExists(request.Owner, request.Repository, request.ReleaseBranch, request.PersonalAccessToken)
	if err != nil {
		http.Error(w, "Error checking if branch exists: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !branchExists {
		http.Error(w, "Branch does not exist in the repository", http.StatusBadRequest)
		return
	}

	baseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s&per_page=100&page=", request.Owner, request.Repository, request.ReleaseBranch)
	method := "GET"

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		http.Error(w, "Error generating new request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	AddhttpAuthRequestHeaders(req, request.PersonalAccessToken)

	commitReleased := false
	var commits []model.CommitData
	client := &http.Client{}

	for page_number := 1; ; page_number++ {
		commits, errMessage = getCommitsByPageNumber(baseUrl, page_number, req, client)
		if errMessage != "" {
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
		if len(commits) == 0 {
			break
		}

		for _, c := range commits {
			if c.SHA == request.CommitId {
				commitReleased = true
				break
			}
		}

		if commitReleased {
			break
		}
	}

	err = json.NewEncoder(w).Encode(model.CommitReleasedResponse{
		CommitReleased: commitReleased,
	})
	if err != nil {
		http.Error(w, "Error while adding json response: "+err.Error(), http.StatusInternalServerError)
	}
}

func getCommitsByPageNumber(baseUrl string, page_number int, req *http.Request, client *http.Client) (commits []model.CommitData, errMessage string) {
	url := baseUrl + strconv.Itoa(page_number)
	u, err := urlpkg.Parse(url)
	if err != nil {
		errMessage = "Error generating new URL: " + err.Error()
		return
	}
	req.URL = u

	res, err := client.Do(req)
	if err != nil {
		errMessage = "Error making request to GitHub: " + err.Error()
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errMessage = "Error reading response: " + err.Error()
		return
	}

	// Unmarshal JSON data into commits variable
	err = json.Unmarshal(body, &commits)
	if err != nil {
		errMessage = "Error unmarshalling JSON: " + err.Error()
	}
	return
}

func (h *commitHandler) GetJobsByCommit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]
	commitSHA := r.URL.Query().Get("commitSHA")
	if commitSHA == "" {
		http.Error(w, "SHA parameter is required", http.StatusBadRequest)
		return
	}

	statuses, err := GetCommitStatuses(owner, repo, commitSHA)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching commit statuses: %v", err), http.StatusInternalServerError)
		return
	}

	statusesJSON, err := json.Marshal(statuses)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling statuses to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(statusesJSON); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
}

func GetCommitStatuses(owner, repo, commitSHA string) ([]Status, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s/status", owner, repo, commitSHA)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var statusResp GitHubStatusResponse
	if err := json.Unmarshal(body, &statusResp); err != nil {
		return nil, err
	}

	return statusResp.Statuses, nil
}
