SHELL := /bin/bash
BUILD_NUMBER := `date +%Y%m%d%H%M`
GREEN := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET := $(shell tput -Txterm sgr0)

.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "${YELLOW}%-16s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download dependencies
	@go mod tidy
	@go mod download

build: ## Build binary for local operating system
	@go generate ./...
	@go build -o mylinter main.go

build-linux: ## Build binary for lnux operating system
	@mkdir -p pkg/linux_amd64/
	@go generate ./...
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o pkg/linux_amd64/mylinter main.go
	@tar -czf pkg/linux_amd64/mylinter-linux_amd64.tar.gz -C pkg/linux_amd64/ mylinter

build-darwin: ## Build binary for darwin operating system
	@mkdir -p pkg/darwin_amd64/
	@go generate ./...
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o pkg/darwin_amd64/mylinter main.go
	@tar -czf pkg/darwin_amd64/mylinter-darwin_amd64.tar.gz -C pkg/darwin_amd64/ mylinter

clean: ## Remove build related file
	@go clean

release: ## Create release
	./release.sh
