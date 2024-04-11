package git_service

import (
	"encoding/json"
	"net/http"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := "Hello World"
	json.NewEncoder(w).Encode(resp)
}
