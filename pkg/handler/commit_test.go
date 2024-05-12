package handler

import (
	"encoding/json"
	"fmt"
	"git-service/pkg/model"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func Test_commitHandler_GetCommitByName(t *testing.T) {
	// Retrieve GitHub personal access token from environment variable
	token := os.Getenv("SARTHAK_GITHUB_PERSONAL_ACCESS_TOKEN")

	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "gcivil-nyu-org", "repo": "INT2-Monday-Spring2024-Team-1"})
	req.URL.RawQuery = fmt.Sprintf("personalAccessToken=%s&branch=master&message=update", token)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	// Call the handler function with the mock request and response
	h.GetCommitByName(rr, req)

	// Check the status code of the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body into a slice of CommitData
	var resp []model.CommitData
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check if the response contains the expected commit message
	expectedMessage := "update"
	foundIssue := false
	for _, c := range resp {
		if !strings.Contains(strings.ToLower(c.Commit.Message), strings.ToLower(expectedMessage)) {
			foundIssue = true
			break
		}
	}
	if foundIssue {
		t.Errorf("Expected response to contain commit message %q, but it was not found", expectedMessage)
	}
}

func Test_commitHandler_GetCommitByName_Branch_Does_Not_Exist(t *testing.T) {
	// Retrieve GitHub personal access token from environment variable
	token := os.Getenv("SARTHAK_GITHUB_PERSONAL_ACCESS_TOKEN")

	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = fmt.Sprintf("branch=testing&personalAccessToken=%s", token)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	h.GetCommitByName(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_commitHandler_CommitReleased_Branch_Does_Not_Exist(t *testing.T) {
	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = "commit_id=1ed709f8ce346c3487cd09eb0875f11efd9bb2dd&release_branch=testing"

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	h.CommitReleased(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_commitHandler_CommitReleased_Commit_Does_Not_Exist(t *testing.T) {
	// Retrieve GitHub personal access token from environment variable
	token := os.Getenv("SARTHAK_GITHUB_PERSONAL_ACCESS_TOKEN")

	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = fmt.Sprintf("commit_id=1ed7hd7sce346c3487cd09eb0875f11efd9bb2dd&release_branch=main&personalAccessToken=%s", token)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	h.CommitReleased(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var resp model.CommitReleasedResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	if resp.CommitReleased {
		t.Errorf("Expected false, but got true")
	}
}

func Test_commitHandler_CommitReleased_Commit_Exists(t *testing.T) {
	// Retrieve GitHub personal access token from environment variable
	token := os.Getenv("SARTHAK_GITHUB_PERSONAL_ACCESS_TOKEN")

	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = fmt.Sprintf("commit_id=ba7daee4b67892dfce920514a3a8fab7fa717fce&release_branch=main&personalAccessToken=%s", token)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	h.CommitReleased(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var resp model.CommitReleasedResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	if !resp.CommitReleased {
		t.Errorf("Expected true, but got false")
	}
}

func TestGetJobsByCommit(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/nyuoss/git-service/commit/getJobsByCommit?commitSHA=4c12ee9e449e8c9dd750783ab00f9be717f2575a", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := mux.NewRouter()
	ch := &commitHandler{}
	r.HandleFunc("/v1/{owner}/{repo}/commit/getJobsByCommit", ch.GetJobsByCommit).Methods("GET")

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

	expected := `[{"url":"https://api.github.com/repos/nyuoss/git-service/statuses/4c12ee9e449e8c9dd750783ab00f9be717f2575a","avatar_url":"https://avatars.githubusercontent.com/in/302869?v=4","id":28952124332,"node_id":"SC_kwDOLf0zQs8AAAAGva5brA","state":"success","description":"","target_url":"https://app.circleci.com/pipelines/circleci/UBj8mMZkjsMXdSRaBgKksF/NefyNZWCQBF8Cojq6e6h7h/40/workflows/e26f6bd8-c521-40d5-9524-20d1b394d8ba","context":"ci/circleci: merge-check","created_at":"2024-05-05T18:06:42Z","updated_at":"2024-05-05T18:06:42Z"}]`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCommitsAfter(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "aryamanrishabh", "repo": "lms-fe"})
	req.URL.RawQuery = "commit=b9e9d948361d97c79bcf0b421d02474e4fda375a"

	rr := httptest.NewRecorder()

	h := &commitHandler{}

	h.GetCommitsAfter(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var resp map[string]map[string][]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned incorrect content type: got %v want %v", contentType, expectedContentType)
	}

	expectedCommit := "42944b800458d6bf85f77c5b6728a1127c4c72d6"
	foundIssue := false
	for _, commitIDs := range resp["commits"] {
		for _, commitID := range commitIDs {
			if commitID == expectedCommit {
				foundIssue = false
				break
			} else {
				foundIssue = true
			}
		}
	}

	if foundIssue {
		t.Errorf("Expected the response to contain commit %s, but it was not found", expectedCommit)
	}
}

func Test_commitHandler_GetCommitByAuthor(t *testing.T) {
	router := mux.NewRouter()
	ch := &commitHandler{}
	router.HandleFunc("/{owner}/{repo}/commits/by-author", ch.GetCommitByAuthor).Methods("GET")

	ts := httptest.NewServer(router)
	defer ts.Close()

	url := fmt.Sprintf("%s/exampleOwner/exampleRepo/commits/by-author?author=JohnDoe&personalAccessToken=testToken", ts.URL)
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	var resp []model.CommitData
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if len(resp) == 0 {
		t.Errorf("Expected at least one commit to be returned, but got none")
	}
}

func Test_commitHandler_GetCommitByAuthor_Author_Does_Not_Exist(t *testing.T) {
	// Retrieve GitHub personal access token from environment variable
	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")

	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = fmt.Sprintf("author=unknownAuthor&personalAccessToken=%s", token)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	// Call the handler function with the mock request and response
	h.GetCommitByAuthor(rr, req)

	// Check the status code of the response
	if rr.Code != http.StatusBadRequest && rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, rr.Code)
	}
}
