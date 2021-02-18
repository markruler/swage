.DEFAULT_GOAL := all
GOOS					:= $(shell go env GOOS)
GOARCH				:= $(shell go env GOARCH)
BIN_DIR				:= bin
BIN_FILE			:= swage
VERSION				:= $(file < ./VERSION)
# GOBIN					:= $(shell go env GOPATH)/bin
# TARGETS				:= darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64

.PHONY: \
				all \
				run \
				version \
				deps \
				gofmt \
				test \
				testv \
				cover \
				clean \
				build \
				docker

all: version deps test cover build

version:
	@echo $(VERSION)

deps:
	@#go get -u
	@#rm -rf vendor/
	@go get -t -d -v ./...
	@go mod tidy
	@go mod vendorbi

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...
	@#golangci-lint linters
	@#golangci-lint run --enable revive
	@#golangci-lint run --enable-all

test: gofmt
	go test ./... --cover

testv: gofmt
	go test ./... -v --cover

cover:
	@scripts/cover --html
	@scripts/cover

run:
	go run main.go

clean:
	go clean
	rm -f *.xlsx
	rm -rf $(BIN_DIR)

build: test
	@GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_DIR)/$(BIN_FILE) main.go

docker:
	@scripts/docker.sh
