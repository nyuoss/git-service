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
			foundIssue = true
		}
	}

	if foundIssue {
		t.Errorf("Expected the response to contain branch %s, but it was not found", expectedBranch)
	}
}

// unit test for checkIfBranchExists function
func Test_checkIfBranchExists(t *testing.T) {
	type args struct {
		owner               string
		repo                string
		branch              string
		personalAccessToken string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
		wantErr    bool
	}{
		{
			name: "Branch exists",
			args: args{
				owner:  "nyuoss",
				repo:   "git-service",
				branch: "main",
			},
			wantExists: true,
			wantErr:    false,
		},
		{
			name: "Branch does not exist",
			args: args{
				owner:  "nyuoss",
				repo:   "git-service",
				branch: "test_main",
			},
			wantExists: false,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, err := checkIfBranchExists(tt.args.owner, tt.args.repo, tt.args.branch, tt.args.personalAccessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfBranchExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExists != tt.wantExists {
				t.Errorf("checkIfBranchExists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}
