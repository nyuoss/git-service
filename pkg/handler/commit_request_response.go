package handler

import (
	"git-service/pkg/model"
	"net/http"

	"github.com/gorilla/mux"
)

// swagger:route GET /{owner}/{repo}/commit/getCommitsBefore commit getCommitsBefore
//
// Provides the commits before the given commit in the branches it is present in.
//
//     Responses:
//       200: GetCommitsBeforeResponse

// swagger:parameters getCommitsBefore
type GetCommitsBeforeReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetCommitsBeforeReqWrapper struct {
	// in:body
	Body GetCommitsBeforeReq `json:"body"`
}

type GetCommitsBeforeResp struct {
	// map[branch-name] ---> [commit1, commit2, commit3...]
	Commits map[string][]string `json:"commits"`
}

// swagger:response GetCommitsBeforeResponse
type GetCommitsBeforeRespWrapper struct {
	// in:body
	Body GetCommitsBeforeResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/commit/getCommitsAfter commit getCommitsAfter
//
// Provides the commits after the given commit in the branches it is present in.
//
//     Responses:
//       200: GetCommitsAfterResponse

// swagger:parameters getCommitsAfter
type GetCommitsAfterReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetCommitsAfterReqWrapper struct {
	// in:body
	Body GetCommitsAfterReq `json:"body"`
}

type GetCommitsAfterResp struct {
	// map[branch-name] ---> [commit1, commit2, commit3...]
	Commits map[string][]string `json:"commits"`
}

// swagger:response GetCommitsAfterResponse
type GetCommitsAfterRespWrapper struct {
	// in:body
	Body GetCommitsAfterResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/commit/getCommitByMessage commit getCommitByMessage
//
// Accepts a message about commit . Provides commit id and name.
//
//     Responses:
//       200: GetCommitByMessageResponse

// swagger:parameters getCommitByMessage
type GetCommitByMessageReq struct {
	// message
	// in: query
	Message string `json:"message"`
}

type GetCommitByMessageReqWrapper struct {
	// in:body
	Body GetCommitByMessageReq `json:"body"`
}

type GetCommitByMessageResp struct {
	// TODO: please add
}

// swagger:response GetCommitsBeforeResponse
type GetCommitByMessageRespWrapper struct {
	// in:body
	Body GetCommitByMessageResp `json:"body"`
}

func GetCommitByMessageRequest(r *http.Request) (req model.GetCommitByMessageRequest, errMessage string) {
	vars := mux.Vars(r)
	queryParams := r.URL.Query()

	owner := vars["owner"]
	if len(owner) == 0 {
		errMessage = "Owner cannot be empty"
		return
	}

	repo := vars["repo"]
	if len(repo) == 0 {
		errMessage = "Repository cannot be empty"
		return
	}

	message := queryParams.Get("message")
	if len(message) == 0 {
		errMessage = "Commit Message cannot be empty"
		return
	}

	req = model.GetCommitByMessageRequest{
		Owner:         owner,
		Repository:    repo,
		CommitMessage: message,
	}
	return
}

func GetCommitReleasedRequest(r *http.Request) (req model.CommitReleasedRequest, errMessage string) {
	vars := mux.Vars(r)
	queryParams := r.URL.Query()

	owner := vars["owner"]
	if len(owner) == 0 {
		errMessage = "Owner cannot be empty"
		return
	}

	repo := vars["repo"]
	if len(repo) == 0 {
		errMessage = "Repository cannot be empty"
		return
	}

	commit_id := queryParams.Get("commit_id")
	if len(commit_id) == 0 {
		errMessage = "Commit Id cannot be empty"
		return
	}

	release_branch := queryParams.Get("release_branch")
	if len(release_branch) == 0 {
		errMessage = "Release Branch cannot be empty"
		return
	}

	req = model.CommitReleasedRequest{
		Owner:         owner,
		Repository:    repo,
		CommitId:      commit_id,
		ReleaseBranch: release_branch,
	}
	return
}
