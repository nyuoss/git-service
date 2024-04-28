package handler

// swagger:route GET /{owner}/{repo}/branch/getActiveBranches branch getActiveBranches
//
// Provides a list of branches that have been active in the given time frame.
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
	// in:body
	Body GetActiveBranchesResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/branch/getBranchByTag branch getBranchByTag
//
// Provides the name of the branch from which the given tag was built.
//
//     Responses:
//       200: GetBranchByTagResponse

// swagger:parameters getBranchByTag
type GetBranchByTagReq struct {
	// tag name
	// in: query
	Tag string `json:"tag"`
}

type GetBranchByTagReqWrapper struct {
	// in:body
	Body GetBranchByTagReq `json:"body"`
}

type GetBranchByTagResp struct {
	Branches string `json:"branches"`
}

// swagger:response GetBranchByTagResponse
type GetBranchByTagRespWrapper struct {
	// in:body
	Body GetBranchByTagResp `json:"body"`
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
