package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type TagHandler interface {
	GetChildTagsByCommit(http.ResponseWriter, *http.Request)
	GetParentTagsByCommit(http.ResponseWriter, *http.Request)
}

var _ TagHandler = &tagHandler{}

type tagHandler struct{}

func NewTagHandler() TagHandler {
	return &tagHandler{}
}

func (h *tagHandler) GetChildTagsByCommit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]

	queryParams := r.URL.Query()

	// Get the value of a specific query parameter
	commitSha := queryParams.Get("commit")

	client := resty.New()

	// get tags
	var repoTags []Tag

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo)
	resp, err := client.R().SetResult(&repoTags).Get(apiURL)
	fmt.Println(repoTags)
	if err != nil {
		http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
		return
	}

	// for fast search commit-sha --> tag
	commitTagMap := make(map[string]string)
	for _, t := range repoTags {
		commitTagMap[t.Commit.SHA] = t.Name
	}

	// get branches
	var repoBranches []Branch
	apiURL = fmt.Sprintf("https://api.github.com/repos/%s/%s/branches", owner, repo)
	resp, err = client.R().SetResult(&repoBranches).Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
		return
	}

	// for each branch, get child tags for the commit
	childTagsByBranch := make(map[string][]string)
	for _, b := range repoBranches {
		sha := b.Commit.SHA
		var commits []Commit
		apiURL = fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s", owner, repo, sha)
		resp, err = client.R().SetResult(&commits).Get(apiURL)
		for _, c := range commits {
			fmt.Println(c.SHA)
		}
		if err != nil {
			http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
			return
		}
		if resp.StatusCode() != http.StatusOK {
			http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
			return
		}

		commitFound := false
		childTags := make([]string, 0)
		for _, c := range commits {
			if commitTagMap[c.SHA] != "" {
				childTags = append(childTags, commitTagMap[c.SHA])
			}
			if strings.HasPrefix(c.SHA, commitSha) {
				commitFound = true
				break
			}
		}

		if !commitFound {
			childTags = []string{}
		}

		childTagsByBranch[b.Name] = childTags
	}

	response := GetChildTagsByCommitResp{
		Tags: childTagsByBranch,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (h *tagHandler) GetParentTagsByCommit(w http.ResponseWriter, r *http.Request) {
	// TODO
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]

	queryParams := r.URL.Query()

	// Get the value of a specific query parameter
	commitSha := queryParams.Get("commit")

	client := resty.New()

	// get tags
	var repoTags []Tag

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", owner, repo)
	resp, err := client.R().SetResult(&repoTags).Get(apiURL)
	fmt.Println(repoTags)
	if err != nil {
		http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
		return
	}

	// for fast search commit-sha --> tag
	commitTagMap := make(map[string]string)
	for _, t := range repoTags {
		commitTagMap[t.Commit.SHA] = t.Name
	}

	// get branches
	var repoBranches []Branch
	apiURL = fmt.Sprintf("https://api.github.com/repos/%s/%s/branches", owner, repo)
	resp, err = client.R().SetResult(&repoBranches).Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
		return
	}

	// for each branch, get child tags for the commit
	childTagsByBranch := make(map[string][]string)
	for _, b := range repoBranches {
		sha := b.Commit.SHA
		var commits []Commit
		apiURL = fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=%s", owner, repo, sha)
		resp, err = client.R().SetResult(&commits).Get(apiURL)
		for _, c := range commits {
			fmt.Println(c.SHA)
		}
		if err != nil {
			http.Error(w, "Failed to fetch branches from the GitHub API", http.StatusInternalServerError)
			return
		}
		if resp.StatusCode() != http.StatusOK {
			http.Error(w, "GitHub API returned status code: %d", resp.StatusCode())
			return
		}

		commitFound := false
		childTags := make([]string, 0)
		for _, c := range commits {
			if commitTagMap[c.SHA] != "" {
				childTags = append(childTags, commitTagMap[c.SHA])
			}
			if strings.HasPrefix(c.SHA, commitSha) {
				commitFound = true
				break
			}
		}

		if !commitFound {
			childTags = []string{}
		}

		childTagsByBranch[b.Name] = childTags
	}

	response := GetChildTagsByCommitResp{
		Tags: childTagsByBranch,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
