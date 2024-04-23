// Package git_service GitsService.
//
// Documentation of our awesome API.
//
//	Schemes: http
//	BasePath: /v1/
//	Version: 1.0.0
//	Host: localhost:8000
//
// swagger:meta
package git_service

const (
	UrlPrefix    = "/v1/{owner}/{repo}"
	BranchPrefix = UrlPrefix + "/branch"
	CommitPrefix = UrlPrefix + "/commit"
	TagPrefix    = UrlPrefix + "/tag"
	JobPrefix    = UrlPrefix + "/job"
)

type Server interface {
	HandleSwagger()

	HandleBranches()
	HandleCommits()
	HandleTags()
	HandleJobs()

	Run()
}
