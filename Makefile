SHELL=/bin/bash
LDFLAGS=-ldflags "-s -w"

CACHE_DIR ?= cache
OUT_DIR ?= out
ASSET_DIR ?= assets

GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

.PHONY: deps
deps:
	go get ${u} -d
	go mod tidy

$(GOBIN)/golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.54.2

.PHONY: test
test:
	go test -v -short -race -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint: $(GOBIN)/golangci-lint
	$(GOBIN)/golangci-lint run

.PHONY: lintfix
lintfix: $(GOBIN)/golangci-lint
	$(GOBIN)/golangci-lint run --fix --timeout 5m

.PHONY: build
build:
	go build $(LDFLAGS) ./cmd/tunnel