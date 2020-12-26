GOOS		:= $(shell go env GOOS)
GOARCH	:= $(shell go env GOARCH)
# GOBIN		:= $(shell go env GOPATH)/bin
BINDIR	:= bin
BINARY	:= swage
VERSION	:= $(file < ./VERSION)
# TARGETS	:= darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64

.PHONY: all
all: version deps test cover run

.PHONY: version
version:
	@echo $(GOBIN)/$(BINARY).$(VERSION)

.PHONY: deps
deps:
	@#go get -u
	@#rm -rf vendor/
	@go get -t -d -v ./...
	@go mod tidy
	@go mod vendor

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: test
test: gofmt
	go test ./... --cover

.PHONY: testv
testv: gofmt
	go test ./... -v --cover

.PHONY: cover
cover:
	@scripts/cover --html
	@scripts/cover

.PHONY: run
run:
	go run main.go gen examples/testdata/json/editor.swagger.json

.PHONY: clean
clean:
	rm -f *.xlsx
	rm -rf $(BINDIR)

.PHONY: build
build:
	@# VERSION := $(cat ./VERSION)
	@GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINARY) main.go

.PHONY: docker
docker:
	@scripts/docker.sh
