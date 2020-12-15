BINARY := swage

.PHONY: test
test:
	@# go test ./... -v
	go test ./cmd/ -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deps
deps:
	@# go get -u
	@go get -t -d -v ./...
	@go mod tidy
	@go mod vendor

# .PHONY: docker
# docker:
# 	./aio/scripts/docker.sh

.PHONY: clean
clean:
	@rm -f ${BINARY}-linux-amd64
	@rm -f ${BINARY}-darwin-amd64
	@rm -f ${BINARY}-windows-amd64

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux-amd64 main.go
	@#GOOS=darwin GOARCH=amd64 go build -o ${BINARY}-darwin-amd64 main.go
	@#GOOS=windows GOARCH=amd64 go build -o ${BINARY}-windows-amd64 main.go

.PHONY: run
run:
	go run main.go aio/example/short.json
