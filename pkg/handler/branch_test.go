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
	req, err := http.NewRequest("GET", "/v1/aryamanrishabh/metricsjs/branch/getActiveBranches?unit=h&number=1000", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := mux.NewRouter()
	h := &branchHandler{}
	r.HandleFunc("/v1/{owner}/{repo}/branch/getActiveBranches", h.GetActiveBranches).Methods("GET")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned incorrect status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned incorrect content type: got %v want %v", contentType, expectedContentType)
	}

	expected := `{"branches":["master"]}`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

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
