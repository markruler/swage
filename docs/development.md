# Development

## Prerequisite

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.37.0
golangci-lint --version
```

## Dependencies

```bash
make deps
```

## Build

```bash
make build
bin/swage version
```

## Run

### on Local Machine

```bash
make run
```

> output >> swage.xlsx

### on Docker

```bash
# Linux, Darwin
make docker

# Windows
aio/scripts/docker.ps1

```

> output >> examples/testdata/docker-swage.xlsx
