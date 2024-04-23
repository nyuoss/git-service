package handler

// swagger:route GET /{owner}/{repo}/job/getJobsByCommit job getJobsByCommit
//
// Accepts a commit id. Provides the CI/CD jobs associated with the given commit.
//
//     Responses:
//       200: GetJobsByCommitResponse

// swagger:parameters getJobsByCommit
type GetJobsByCommitReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetJobsByCommitReqWrapper struct {
	// in:body
	Body GetJobsByCommitReq `json:"body"`
}

type GetJobsByCommitResp struct {
	// [jobUrl1, jobUrl2, jobUrl3...]
	Jobs []string `json:"jobs"`
}

// swagger:response GetJobsByCommitResponse
type GetJobsByCommitRespWrapper struct {
	// in:body
	Body GetJobsByCommitResp `json:"body"`
}
