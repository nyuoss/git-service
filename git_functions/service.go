package git_service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := "Hello World"

	_ = json.NewEncoder(w).Encode(resp)
}

func GetActiveBranches(w http.ResponseWriter, r *http.Request) {
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
