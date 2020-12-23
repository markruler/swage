BINARY = swage
VERSION = $(file < ./VERSION)

.PHONY: all
all: echo deps test cover run

.PHONY: echo
echo:
	@echo ${BINARY}.${VERSION}

.PHONY: deps
deps:
	@rm -rf vendor/
	@# go get -u
	@go get -t -u -d -v ./...
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
	@aio/scripts/cover --html
	@aio/scripts/cover

.PHONY: run
run:
	go run main.go gen aio/example/short.json

.PHONY: clean
clean:
	@rm -f *.xlsx
	@rm -f ${BINARY}.${VERSION}-linux-amd64
	@rm -f ${BINARY}.${VERSION}-darwin-amd64
	@rm -f ${BINARY}.${VERSION}-windows-amd64

.PHONY: build
build:
	@# VERSION := $(cat ./VERSION)
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY}.${VERSION}-linux-amd64 main.go
	@#GOOS=darwin GOARCH=amd64 go build -o ${BINARY}.${VERSION}-darwin-amd64 main.go
	@#GOOS=windows GOARCH=amd64 go build -o ${BINARY}.${VERSION}-windows-amd64 main.go

.PHONY: docker
docker:
	@aio/scripts/docker.sh
