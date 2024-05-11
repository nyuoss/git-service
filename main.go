package main

import (
	gitService "git-service/git_functions"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	server := gitService.NewGitServer(router)

	// serve swagger-ui
	server.HandleSwagger()

	server.HandleBranches()
	server.HandleCommits()
	server.HandleJobs()
	server.HandleTags()

	server.Run()
}
