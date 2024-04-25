package handler

// swagger:parameters getActiveBranches getBranchByTag getChildTagsByCommit getParentTagsByCommit getCommitsBefore getCommitsAfter getCommitByMessage getJobsByCommit
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
