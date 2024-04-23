package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BranchHandler interface {
	GetActiveBranches(http.ResponseWriter, *http.Request)
	GetBranchByTag(http.ResponseWriter, *http.Request)
}

var _ BranchHandler = &branchHandler{}

type branchHandler struct{}

func NewBranchHandler() BranchHandler {
	return &branchHandler{}
}

func (h *branchHandler) GetActiveBranches(w http.ResponseWriter, r *http.Request) {
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

func (h *branchHandler) GetBranchByTag(w http.ResponseWriter, r *http.Request) {}
