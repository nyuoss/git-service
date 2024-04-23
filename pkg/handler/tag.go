package handler

import "net/http"

type TagHandler interface {
	GetChildTagsByCommit(http.ResponseWriter, *http.Request)
	GetParentTagsByCommit(http.ResponseWriter, *http.Request)
}

var _ TagHandler = &tagHandler{}

type tagHandler struct{}

func NewTagHandler() TagHandler {
	return &tagHandler{}
}

func (h *tagHandler) GetChildTagsByCommit(w http.ResponseWriter, r *http.Request) {}

func (h *tagHandler) GetParentTagsByCommit(w http.ResponseWriter, r *http.Request) {}
