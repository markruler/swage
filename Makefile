BINARY = swage
VERSION = $(file < ./VERSION)

.PHONY: echo
echo:
	echo ${BINARY}.${VERSION}

.PHONY: test
test:
	@# go test ./... -v
	go test ./pkg/... -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deps
deps:
	@# go get -u
	@go get -t -d -v ./...
	@go mod tidy
	@go mod vendor

.PHONY: docker
docker:
	@aio/scripts/docker.sh

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

.PHONY: run
run:
	go run main.go gen aio/example/short.json
