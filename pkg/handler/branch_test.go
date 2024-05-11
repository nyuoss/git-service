package handler

import "testing"

// unit test for checkIfBranchExists function
func Test_checkIfBranchExists(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		branch string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
		wantErr    bool
	}{
		{
			name: "Branch exists",
			args: args{
				owner:  "nyuoss",
				repo:   "git-service",
				branch: "main",
			},
			wantExists: true,
			wantErr:    false,
		},
		{
			name: "Branch does not exist",
			args: args{
				owner:  "nyuoss",
				repo:   "git-service",
				branch: "test_main",
			},
			wantExists: false,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, err := checkIfBranchExists(tt.args.owner, tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfBranchExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExists != tt.wantExists {
				t.Errorf("checkIfBranchExists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}
