# Git-Service: Enhanced Git Navigation Tool

## Purpose
It is hard for developers to track branches and tags in GitHub. Especially in an active work group where developers collaborate and work on multiple branches, cherry-pick, merge, rebase and patch operations make it even more difficult to search for a commit. Searching for a commit may be required in cases like code review, debugging, reverting specific changes, etc. Moreover, lots of continuous integration jobs are triggered by events or run periodically every day. Hence, finding out relationships between commits and jobs is also a tough task. This project is crucial because it addresses these challenges by introducing a tool that simplifies navigating through commits, branches, tags, and clarifies the connections between the code changes and CI/CD processes.

## Features
1. `/<owner>/<repo>/branch/getActiveBranches`: Provides a list of branches that have been active in the given time frame
2. `/<owner>/<repo>/branch/getBranchByTag`: Provides the name of the branch from which the given tag was built
3. `/<owner>/<repo>/tag/getChildTagsByCommit`: Provides the child tags in each branch
4. `/<owner>/<repo>/tag/getParentTagsByCommit`: Provides the nearest parent tags in each branch
5. `/<owner>/<repo>/commit/getCommitsBefore`: Provides the commits before the given commit in the branches it is present in
6. `/<owner>/<repo>/commit/getCommitsAfter`: Provides the commits after the given commit in the branches it is present in
7. `/<owner>/<repo>/commit/getCommitByName`: Provides list of commits when searched by name
8. `/<owner>/<repo>/commit/getCommitByAuthor`: Provides list of commits when searched by author
9. `/<owner>/<repo>job/getJobsByCommit`: Provides the CI/CD jobs associated with the given commit
10. `/<owner>/<repo>/commit/commitReleased`: Accepts commit id and release branch name. Tells if the commit is released or not


## Prerequisites: Installation Instructions
- Install the `golang` compiler from the [official source](https://go.dev) (version 1.21.3)
- Include $GOPATH in your path via `export PATH=$PATH:$(go env GOPATH)/bin`
- Adding the above to your `~/.bashrc` file will require `source ~/.bashrc`
- Alternatively, if using a package manager like homebrew for example, install golang by running `brew install golang`

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
|____ git_functions	# Git-Service Server Component
|____Makefile       # make test, fmt….
|____README.md
|____.gitignore
|____.circleci	    # CI configuration
|____main.go	    # main function
|____pkg            # Components
|   ├── handler     # Server Handler Components
|   ├── model       # Github-API defined Objects
|   └── swagger-ui  # Swagger File and Frontend
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
1. Clone the repository
2. Pull latest code on the `main` branch
3. Follow the instructions given in the `Usage Instructions: Getting Started` section to start your server
4. Swagger UI should be accessible on this link: http://localhost:8000/swaggerui/

### Adding New APIs: UI
1. For the endpoint to be accessible on UI, make changes to `swagger.json` file. You can pick any of the existing endpoints and it is fairly straightforward to follow.
2. After making changes to this file, hard reload your web browser for the changes to reflect, in case the previous state is cached.

### Adding New APIs: Backend
1. The `main.go` file initializes all packages.
2. For example, `HandleCommits` is defined in the `git_functions/service.go` file.
3. It states all endpoints related to commits on GitHub.
4. To define a new endpoint, simply add it here.
5. The corresponding function call is defined in the `pkg/handler/commit.go` file.
6. You can add your function and the respective implementation here.
7. If you were to add an endpoint related to GitHub branches, then defining it in `pkg/handler/branch.go` would be more appropriate.
8. After adding the backend code, you will have to run `go build` and `go run main.go` to test your changes.


### Testing Endpoints
1. Pick the API you wish to test from Swagger UI.
2. Click on `Try it out` and fill the request parameters.
3. Click on `Execute` to make the backend call.
4. To debug backend code, add `fmt.Println` statements at suitable breakpoints.

### Contribution Templates
- [Pull Request Template](./pull_request_template.md)
- [Issue Template](./issue_template.md)
