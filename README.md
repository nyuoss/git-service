# Git-Service: Enhanced Git Navigation Tool

## Purpose
It is hard for developers to track branches and tags in GitHub. Especially in an active work group where developers collaborate and work on multiple branches, cherry-pick, merge, rebase and patch operations make it even more difficult to search for a commit. Searching for a commit may be required in cases like code review, debugging, reverting specific changes, etc. Moreover, lots of continuous integration jobs are triggered by events or run periodically every day. Hence, finding out relationships between commits and jobs is also a tough task. This project is crucial because it addresses these challenges by introducing a tool that simplifies navigating through commits, branches, tags, and clarifies the connections between the code changes and CI/CD processes.

## Features
1. `/v1/<owner>/<github-repo-name>/branch/getActiveBranches`: Provides a list of branches that have been active in the given time frame
2. `/v1/<owner>/<github-repo-name>/branch/getBranchByTag`: Provides the name of the branch from which the given tag was built
3. `/v1/<owner>/<github-repo-name>/tag/getChildTagsByCommit`: Provides the child tags in each branch
4. `/v1/<owner>/<github-repo-name>/tag/getParentTagsByCommit`: Provides the nearest parent tags in each branch
5. `/v1/<owner>/<github-repo-name>/commit/getCommitsBefore`: Provides the commits before the given commit in the branches it is present in
6. `/v1/<owner>/<github-repo-name>/commit/getCommitsAfter`: Provides the commits after the given commit in the branches it is present in
7. `/v1/<owner>/<github-repo-name>/commit/getCommitByName`: Provides list of commits when searched by name
8. `/v1/<owner>/<github-repo-name>/commit/getCommitByDescription`: Provides list of commits when searched by description
9. `/v1/<owner>/<github-repo-name>/commit/getCommitByAuthor`: Provides list of commits when searched by author
10. `/v1/<owner>/<github-repo-name>job/getJobsByCommit`: Provides the CI/CD jobs associated with the given commit


## Prerequisites: Installation Instructions
* Install the `golang` compiler from the [official source](https://go.dev) (version 1.21.3)
* Include $GOPATH in your path via `export PATH=$PATH:$(go env GOPATH)/bin`
* Adding the above to your `~/.bashrc` file will require `source ~/.bashrc`

### Tools
1. Language: Golang@1.21.3
2. Package Management Tool: go mod
3. Code Format Tool: golangci-lint@1.55.2
4. Static Analysis Tool: golangci-lint@1.55.2
5. CI Tool: CircleCI@2.1

### Structure

``` shell
$ tree
.
|____go.mod         # dependencies
|____LICENSE		
|____Makefile       # make test, fmt….
|____README.md
|____.gitignore
|____.circleci	    # CI configuration
|____main.go	    # main function
|____pkg            # Components
|____.golangci.yml  # golangci-lint Configuration
```

## Usage Instructions: Getting Started

You have to prepare tools with make -

``` shell
make tools
```

Also, you can run lint, format and test with make before making a pull request -

``` shell
# test 
make test

# lint
make lint

# format
make format
```

### Start Your Server
``` shell
# Download the packages
go mod download

# Compile the packages and dependencies
go build

# Run the code
go run main.go
```

Use http://localhost:8000/swaggerui/ in your browser to test the APIs

## Contribution Guidelines
+ [pull_request_template.md](./pull_request_template.md)
+ [issue_template.md](./issue_template.md)
