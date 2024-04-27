package handler

import (
    "encoding/json"
    "fmt"
	"io"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/go-resty/resty/v2"
)
type commit struct {
    SHA     string `json:"sha"`
    Commit  struct {
        Author struct {
            Date string `json:"date"`
        } `json:"author"`
    } `json:"commit"`
}

type CommitHandler interface {
	GetCommitsBefore(http.ResponseWriter, *http.Request)
	GetCommitsAfter(http.ResponseWriter, *http.Request)
	GetCommitByMessage(http.ResponseWriter, *http.Request)
}

var _ CommitHandler = &commitHandler{}

type commitHandler struct{}

func NewCommitHandler() CommitHandler {
	return &commitHandler{}
}

func (h *commitHandler) GetCommitsBefore(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    owner := vars["owner"]
    repo := vars["repo"]
    commitID := vars["commit"]

    client := resty.New()
    apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s", owner, repo, commitID)

    resp, err := client.R().
   
        Get(apiUrl)

    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to fetch data: %s", err.Error()), http.StatusInternalServerError)
        return
    }

    if resp.StatusCode() != http.StatusOK {
        body, _ := io.ReadAll(resp.RawResponse.Body)
        http.Error(w, fmt.Sprintf("GitHub API returned status code: %d, body: %s", resp.StatusCode(), string(body)), resp.StatusCode())
        return
    }

    var commits []commit
    if err := json.Unmarshal(resp.Body(), &commits); err != nil {
        http.Error(w, "Failed to parse commits data", http.StatusInternalServerError)
        return
    }

    commitsByBranch := mapCommitsToBranches(commits)
    if err := json.NewEncoder(w).Encode(commitsByBranch); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

func mapCommitsToBranches(commits []commit) map[string][]string {
    result := make(map[string][]string)
    for _, c := range commits {
        branchName := "main"  // Placeholder: you would need to determine the branch somehow
        result[branchName] = append(result[branchName], c.SHA)
    }
    return result
}

func (h *commitHandler) GetCommitsAfter(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *commitHandler) GetCommitByMessage(w http.ResponseWriter, r *http.Request) {
	// TODO
}
