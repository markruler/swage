SHELL := bash
.ONESHELL:

NAME					:= swage
BIN_DIR				:= bin
GOOS					:= $(shell go env GOOS)
GOARCH				:= $(shell go env GOARCH)
VERSION				:= $(file < ./VERSION)
# GOBIN					:= $(shell go env GOPATH)/bin
# TARGETS				:= darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64

.SHELLFLAGS 	:= -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS 		+= --warn-undefined-variables
# go tool link (https://golang.org/cmd/link/)
LDFLAGS 			:= -s -w -extldflags='-static' \
								-X 'github.com/cxsu/swage/pkg/cmd.swageVersion=${VERSION}'

all: version deps test cover build
.PHONY: all

version:
	@echo $(VERSION)
.PHONY: version

tag:
	git tag $(VERSION)
.PHONY: tag

deps:
	@#go get -u
	@#rm -rf vendor/
	@go get -t -d -v ./...
	@go mod tidy
	@go mod vendorbi
.PHONY: deps

fmt:
	go fmt ./...
.PHONY: fmt

lint:
	golangci-lint run ./...
	@#golangci-lint linters
	@#golangci-lint run --enable revive
	@#golangci-lint run --enable-all
.PHONY: lint

test: fmt
	go test ./... --cover
.PHONY: test

testv: fmt
	go test ./... -v --cover
.PHONY: testv

cover:
	@scripts/cover --html
	@scripts/cover
.PHONY: cover

run:
	go run main.go
.PHONY: run

clean:
	go clean
	rm -f *.xlsx
	rm -rf $(BIN_DIR)
.PHONY: clean

build: test
	@GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(NAME) -v -ldflags="${LDFLAGS}" main.go
.PHONY: build

docker:
	@scripts/docker.sh
.PHONY: docker
