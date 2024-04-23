package handler

import "net/http"

type JobHandler interface {
	GetJobsByCommit(http.ResponseWriter, *http.Request)
}

var _ JobHandler = &jobHandler{}

type jobHandler struct{}

func NewJobHandler() JobHandler {
	return &jobHandler{}
}

func (h *jobHandler) GetJobsByCommit(w http.ResponseWriter, r *http.Request) {}
