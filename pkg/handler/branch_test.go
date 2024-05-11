package handler

import (
	"encoding/json"
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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	expectedResponseBody := `{"branches":["master"]}`
	actualResponseBody := strings.TrimSpace(rr.Body.String())

	if actualResponseBody != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", actualResponseBody, expectedResponseBody)
	}
}

func TestGetActiveBranches(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "aryamanrishabh", "repo": "metricsjs"})
	req.URL.RawQuery = "unit=h&number=6000"
	

	rr := httptest.NewRecorder()

	b := &branchHandler{}

	b.GetActiveBranches(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var resp map[string][]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned incorrect content type: got %v want %v", contentType, expectedContentType)
	}

	expectedBranch := "master"
	foundIssue := false
	for _, b := range resp["branches"] {
		if b == expectedBranch {
			foundIssue = false
			break
		} else {
			foundIssue =  true
		}
	}

	if foundIssue {
		t.Errorf("Expected the response to contain branch %s, but it was not found", expectedBranch)
	}
}
