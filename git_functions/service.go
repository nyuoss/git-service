package git_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := "Hello World"

	_ = json.NewEncoder(w).Encode(resp)
}

func GetActiveBranches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]

	queryParams := r.URL.Query()

	// Get the value of a specific query parameter
	unit := queryParams.Get("unit")
	number := queryParams.Get("number")

	branches := []string{owner, repo, unit, number}
	resp := GetActiveBranchesResp{
		Branches: branches,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func GetBranchByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]

	queryParams := r.URL.Query()
	tag := queryParams.Get("tag")

	tags, err := getTags(owner, repo)
	if err != nil {
		http.Error(w, "Failed to get tags from GitHub API", http.StatusInternalServerError)
		return
	}

	var tagMatch bool
	var commitSHA string
	for _, t := range tags {
		if t.Name == tag {
			tagMatch = true
			commitSHA = t.Commit.SHA
			break
		}
	}

	if !tagMatch {
		http.Error(w, "Tag not found", http.StatusNotFound)
		return
	}

	branches, err := getBranches(owner, repo, commitSHA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string][]string{"branches": branches}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getTags(owner, repo string) ([]Tag, error) {
	var tags []Tag
	client := resty.New()

	resp, err := client.R().
		SetResult(&tags).
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status code: %d", resp.StatusCode())
	}

	return tags, nil
}

func getBranches(owner, repo, commitSHA string) ([]string, error) {
	var branches []Branch
	client := resty.New()

	resp, err := client.R().
		SetResult(&branches).
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s/branches-where-head", owner, repo, commitSHA))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status code: %d", resp.StatusCode())
	}

	var branchNames []string
	for _, branch := range branches {
		branchNames = append(branchNames, branch.Name)
	}

	return branchNames, nil
}
