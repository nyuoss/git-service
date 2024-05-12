package handler

import (
	"encoding/json"
	"git-service/pkg/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func Test_commitHandler_GetCommitByMessage(t *testing.T) {
	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "gcivil-nyu-org", "repo": "INT2-Monday-Spring2024-Team-1"})
	req.URL.RawQuery = "message=update allowed hosts again"

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	// Call the handler function with the mock request and response
	h.GetCommitByMessage(rr, req)

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
	expectedMessage := "update allowed hosts again"
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
	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = "commit_id=1ed7hd7sce346c3487cd09eb0875f11efd9bb2dd&release_branch=main"

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
	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = "commit_id=ba7daee4b67892dfce920514a3a8fab7fa717fce&release_branch=main"

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

func Test_commitHandler_GetCommitByAuthor(t *testing.T) {
	// Mock request data
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"owner": "nyuoss", "repo": "git-service"})
	req.URL.RawQuery = "author=exampleUser" // Removed token from query

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock commit handler
	h := &commitHandler{}

	// Call the handler function with the mock request and response
	h.GetCommitByAuthor(rr, req)

	// Check the status code of the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body into a slice of CommitData
	var resp []model.CommitData
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check if the response contains commits by the expected author
	// This part assumes that the GitHub username "exampleUser" should match the author of some commits
	expectedAuthor := "sarthakgoel1997"
	found := false
	for _, c := range resp {
		if c.Author.Login == expectedAuthor {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected at least one commit from author %q, but none were found", expectedAuthor)
	}
}
