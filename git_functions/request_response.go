package git_service

import (
	"go-template/model"
	"net/http"
)

func GetCommitByMessageRequest(w http.ResponseWriter, r *http.Request) (req model.GetCommitByMessageRequest, errMessage string) {
	owner := r.URL.Query().Get("owner")
	if len(owner) == 0 {
		errMessage = "Owner cannot be empty"
		return
	}

	repository := r.URL.Query().Get("repository")
	if len(repository) == 0 {
		errMessage = "Repository cannot be empty"
		return
	}

	personal_access_token := r.URL.Query().Get("personal_access_token")
	if len(personal_access_token) == 0 {
		errMessage = "Personal Access Token cannot be empty"
		return
	}

	commit_message := r.URL.Query().Get("commit_message")
	if len(commit_message) == 0 {
		errMessage = "Commit Message cannot be empty"
		return
	}

	req = model.GetCommitByMessageRequest{
		Owner:               owner,
		Repository:          repository,
		PersonalAccessToken: personal_access_token,
		CommitMessage:       commit_message,
	}
	return
}

func AddRequestHeaders(req *http.Request, personalAccessToken string) {
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+personalAccessToken)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
}
