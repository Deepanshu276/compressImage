# Ensure go modules are enabled:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Disable CGO so that we always generate static binaries:
export CGO_ENABLED=0

# Constants:
GOPATH := $(shell go env GOPATH)

.PHONY: build
build:
	go build -o imgcompress

.PHONY: install
install:
	go build -o ${GOPATH}/bin/imgcompress

.PHONY: mod
mod:
	go mod tidy


.PHONY: test
test: 
	go test ./test -v

.PHONY: lint
lint: getlint
	$(GOPATH)/bin/golangci-lint run

