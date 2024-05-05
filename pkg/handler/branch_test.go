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
	req, err := http.NewRequest("GET", "/v1/aryamanrishabh/metricsjs/branch/getActiveBranches?unit=h&number=300", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := mux.NewRouter()
	h := &branchHandler{}
	r.HandleFunc("/v1/{owner}/{repo}/branch/getActiveBranches", h.GetActiveBranches).Methods("GET")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned incorrect status code")
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned incorrect content type")
	}

	expected := `{"branches":["master"]}`
	got := strings.TrimSpace(rr.Body.String())
	
	if got != expected {
		t.Errorf("Handler returned unexpected body")
	}
}

