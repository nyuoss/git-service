# Git-Service

Git Service: Enhanced Git Navigation Tool

## Prerequisites

* Install the `golang` compiler from the [official source](https://go.dev) (version 1.21.3)
* Include $GOPATH in your path via `export PATH=$PATH:$(go env GOPATH)/bin`
* Adding the above to your `~/.bashrc` file will require `source ~/.bashrc`

### Running the code
``` shell
go run main.go
```

## Introduction

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

## Getting Started

You have to prepare tools with make.

``` shell
make tools
```

Also, you can run lint, format and test with make before making a pull request.

``` shell
# test 
make test

# lint
make lint

# format
make format
```

### Start your server
``` shell
# generate swagger file 
make swagger

go run main.go
```

use http://localhost:8000/swaggerui/ in your browse to test your Api.

## Contributing
+ [pull_request_template.md](./pull_request_template.md)
+ [issue_template.md](./issue_template.md)
