package git_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"git-service/pkg/handler"
)

var _ Server = &gitServer{}

type gitServer struct {
	r *mux.Router

	bh handler.BranchHandler
	ch handler.CommitHandler
	th handler.TagHandler
	jh handler.JobHandler
}

func NewGitServer(r *mux.Router) Server {
	return &gitServer{
		r:  r,
		bh: handler.NewBranchHandler(),
		ch: handler.NewCommitHandler(),
		th: handler.NewTagHandler(),
		jh: handler.NewJobHandler(),
	}
}

func (s *gitServer) Run() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(s.r)

	port := 8000
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func (s *gitServer) HandleSwagger() {
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./pkg/swagger-ui")))
	s.r.PathPrefix("/swaggerui/").Handler(sh)
}

func (s *gitServer) HandleBranches() {
	s.r.HandleFunc(
		BranchPrefix+"/getActiveBranches", s.bh.GetActiveBranches).
		Methods(http.MethodGet)

	s.r.HandleFunc(
		BranchPrefix+"/getBranchByTag", s.bh.GetBranchByTag).
		Methods(http.MethodGet)
}

func (s *gitServer) HandleCommits() {
	s.r.HandleFunc(
		CommitPrefix+"/getCommitByMessage", s.ch.GetCommitByMessage).
		Methods(http.MethodGet)

	s.r.HandleFunc(
		CommitPrefix+"/getCommitsBefore", s.ch.GetCommitsBefore).
		Methods(http.MethodGet)

	s.r.HandleFunc(
		CommitPrefix+"/getCommitsAfter", s.ch.GetCommitsAfter).
		Methods(http.MethodGet)

}

func (s *gitServer) HandleTags() {
	s.r.HandleFunc(
		TagPrefix+"/getChildTagsByCommit", s.th.GetChildTagsByCommit).
		Methods(http.MethodGet)

	s.r.HandleFunc(
		TagPrefix+"/getParentTagsByCommit", s.th.GetParentTagsByCommit).
		Methods(http.MethodGet)
}

func (s *gitServer) HandleJobs() {
	s.r.HandleFunc(
		JobPrefix+"/getJobsByCommit", s.jh.GetJobsByCommit).
		Methods(http.MethodGet)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
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
