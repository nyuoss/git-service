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
	if !foundIssue {
		t.Errorf("Expected response to contain commit message %q, but it was not found", expectedMessage)
	}
}
