package handler

// swagger:route GET /{owner}/{repo}/tag/getChildTagsByCommit tag getChildTagsByCommit
//
// Accepts a commit id as a query parameter. Provides the child tags in each branch.
//
//     Responses:
//       200: GetChildTagsByCommitResponse

// swagger:parameters getChildTagsByCommit
type GetChildTagsByCommitReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetChildTagsByCommitReqWrapper struct {
	// in:body
	Body GetChildTagsByCommitReq `json:"body"`
}

type GetChildTagsByCommitResp struct {
	// map[branch-name] ---> [tag1, tag2, tag3...]
	Tags map[string][]string `json:"Tags"`
}

// swagger:response GetChildTagsByCommitResponse
type GetChildTagsByCommitRespWrapper struct {
	// in:body
	Body GetChildTagsByCommitResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/tag/getParentTagsByCommit tag getParentTagsByCommit
//
// Accepts a commit id as a query parameter. Provides the nearest parent tags in each branch.
//
//     Responses:
//       200: GetParentTagsByCommitResponse

// swagger:parameters getParentTagsByCommit
type GetParentTagsByCommitReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetParentTagsByCommitReqWrapper struct {
	// in:body
	Body GetParentTagsByCommitReq `json:"body"`
}

type GetParentTagsByCommitResp struct {
	// map[branch-name] ---> tag-name
	Tags map[string]string `json:"Tags"`
}

// swagger:response GetParentTagsByCommitResponse
type GetParentTagsByCommitRespWrapper struct {
	// in:body
	Body GetParentTagsByCommitResp `json:"body"`
}
