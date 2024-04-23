package git_service

// swagger:parameters getActiveBranches
type RepoInfo struct {
	// Owner of the repository
	//
	// in: path
	// required: true
	Owner string `json:"owner"`

	// repository name
	//
	// in: path
	// required: true
	Repo string `json:"repo"`
}

// swagger:route GET /{owner}/{repo}/branch/getActiveBranches branch getActiveBranches
//
// Get active branches in a period of time.
//
// This will return a user based on the ID provided in the path parameter.
//
//     Responses:
//       200: GetActiveBranchesResponse

// swagger:parameters getActiveBranches
type GetActiveBranchesReq struct {
	// time unit
	// in: query
	Unit string `json:"unit"`
	// number
	// in: query
	Number int `json:"number"`
}

type GetActiveBranchesReqWrapper struct {
	// in:body
	Body GetActiveBranchesReq `json:"body"`
}

type GetActiveBranchesResp struct {
	Branches []string `json:"branches"`
}

// swagger:response GetActiveBranchesResponse
type GetActiveBranchesRespWrapper struct {
	// Body GetActiveBranchesResponseBody
	// in:body
	Body GetActiveBranchesResp `json:"body"`
}

type Tag struct {
	Name   string `json:"name"`
	Commit struct {
		SHA string `json:"sha"`
	} `json:"commit"`
}

type Branch struct {
	Name      string `json:"name"`
	Commit    Commit `json:"commit"`
	Protected bool   `json:"protected"`
}

type Commit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}
