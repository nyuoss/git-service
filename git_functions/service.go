package git_service

import (
	"fmt"
	"log"
	"net/http"

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
