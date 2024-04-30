package handler

import (
	"encoding/json"
	"fmt"
	"git-service/pkg/model"
	"io"
	"net/http"
	"strconv"
	"strings"

	urlpkg "net/url"
)

type CommitHandler interface {
	GetCommitsBefore(http.ResponseWriter, *http.Request)
	GetCommitsAfter(http.ResponseWriter, *http.Request)
	GetCommitByMessage(http.ResponseWriter, *http.Request)
}

var _ CommitHandler = &commitHandler{}

type commitHandler struct{}

func NewCommitHandler() CommitHandler {
	return &commitHandler{}
}

func (h *commitHandler) GetCommitsBefore(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *commitHandler) GetCommitsAfter(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *commitHandler) GetCommitByMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get request data from query params
	request, errMessage := GetCommitByMessageRequest(r)
	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	baseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?per_page=100&page=", request.Owner, request.Repository)
	method := "GET"

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		http.Error(w, "Error generating new request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := []model.CommitData{}

	for page_number := 1; ; page_number++ {
		url := baseUrl + strconv.Itoa(page_number)
		u, err := urlpkg.Parse(url)
		if err != nil {
			http.Error(w, "Error generating new URL: "+err.Error(), http.StatusInternalServerError)
			return
		}
		req.URL = u

		// Define a variable of type []Commit to store the data
		var commits []model.CommitData

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error making request to GitHub: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Error reading response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal JSON data into commits variable
		err = json.Unmarshal(body, &commits)
		if err != nil {
			http.Error(w, "Error unmarshalling JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if len(commits) == 0 {
			break
		}

		for _, c := range commits {
			if strings.Contains(strings.ToLower(c.Commit.Message), strings.ToLower(request.CommitMessage)) {
				resp = append(resp, c)
			}
		}
	}

	_ = json.NewEncoder(w).Encode(resp)
}
