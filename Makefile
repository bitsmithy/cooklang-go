## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## setup: bootstrap all dependencies for the project
.PHONY: setup
setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest

## audit: tidy dependencies and format, vet, and test all code
.PHONY: audit
audit: format lint test

## format: format codebase using `go fmt`
.PHONY: format
format:
	@echo 'Formatting code...'
	go fmt ./...

## lint: lint and vet codebase using `go vet`, `staticcheck` and `golangci-lint`
.PHONY: lint
lint:
	@echo 'Linting and vetting code...'
	go vet ./...
	staticcheck ./...
	golangci-lint run

## test: test codebase using `gotestsum`
.PHONY: test
test: deps
	@echo 'Running tests...'
	go clean -testcache
	gotestsum --format dots-v2 -- -race -vet=off ./...

## deps: sync dependencies using `go mod tidy` and `go mod verify`
.PHONY: deps
deps:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
