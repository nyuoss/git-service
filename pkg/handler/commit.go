package handler

import (
	"net/http"
)

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

func (h *commitHandler) GetCommitsBefore(w http.ResponseWriter, r *http.Request) {}

func (h *commitHandler) GetCommitsAfter(w http.ResponseWriter, r *http.Request) {}

func (h *commitHandler) GetCommitByMessage(w http.ResponseWriter, r *http.Request) {}
