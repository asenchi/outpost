# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Setup name variables for the package/tool
NAME := outpost
PKG := github.com/infrastellar/$(NAME)

all: clean build fmt lint test staticcheck vet install

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME): *.go
	@echo "+ $@"
	go build -o $(NAME) .

.PHONY: fmt
fmt: ## Verifies all files have men `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test -v $(shell go list ./... | grep -v vendor)

.PHONY: install
install: ## Installs the executable or package
	@echo "+ $@"
	go install -a .

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(RM) $(NAME)
