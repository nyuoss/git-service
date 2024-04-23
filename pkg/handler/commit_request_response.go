package handler

// swagger:route GET /{owner}/{repo}/commit/getCommitsBefore commit getCommitsBefore
//
// Provides the commits before the given commit in the branches it is present in.
//
//     Responses:
//       200: GetCommitsBeforeResponse

// swagger:parameters getCommitsBefore
type GetCommitsBeforeReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetCommitsBeforeReqWrapper struct {
	// in:body
	Body GetCommitsBeforeReq `json:"body"`
}

type GetCommitsBeforeResp struct {
	// map[branch-name] ---> [commit1, commit2, commit3...]
	Commits map[string][]string `json:"commits"`
}

// swagger:response GetCommitsBeforeResponse
type GetCommitsBeforeRespWrapper struct {
	// in:body
	Body GetCommitsBeforeResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/commit/getCommitsAfter commit getCommitsAfter
//
// Provides the commits after the given commit in the branches it is present in.
//
//     Responses:
//       200: GetCommitsAfterResponse

// swagger:parameters getCommitsAfter
type GetCommitsAfterReq struct {
	// commit id
	// in: query
	Commit string `json:"commit"`
}

type GetCommitsAfterReqWrapper struct {
	// in:body
	Body GetCommitsAfterReq `json:"body"`
}

type GetCommitsAfterResp struct {
	// map[branch-name] ---> [commit1, commit2, commit3...]
	Commits map[string][]string `json:"commits"`
}

// swagger:response GetCommitsAfterResponse
type GetCommitsAfterRespWrapper struct {
	// in:body
	Body GetCommitsAfterResp `json:"body"`
}

// swagger:route GET /{owner}/{repo}/commit/getCommitByMessage commit getCommitByMessage
//
// Accepts a message about commit . Provides commit id and name.
//
//     Responses:
//       200: GetCommitByMessageResponse

// swagger:parameters getCommitByMessage
type GetCommitByMessageReq struct {
	// message
	// in: query
	Message string `json:"message"`
}

type GetCommitByMessageReqWrapper struct {
	// in:body
	Body GetCommitByMessageReq `json:"body"`
}

type GetCommitByMessageResp struct {
	// TODO: please add
}

// swagger:response GetCommitsBeforeResponse
type GetCommitByMessageRespWrapper struct {
	// in:body
	Body GetCommitByMessageResp `json:"body"`
}
