package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetBranchByTag(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/Siddharth-Bhardwaj/dining-concierge/branch/getBranchByTag?tag=1.0.0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/{owner}/{repo}/branch/getBranchByTag", NewBranchHandler().GetBranchByTag).Methods("GET")

	router.ServeHTTP(rr, req)

	expectedResponseBody := `{"branches":["master"]}`
	actualResponseBody := strings.TrimSpace(rr.Body.String())

	if actualResponseBody != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", actualResponseBody, expectedResponseBody)
	}
}
