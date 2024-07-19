GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --tags --abbrev=0)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
########################################################################################
default: install

.PHONY: create_bin_dir
create_bin_dir:
	@mkdir -p bin

.PHONY: build
build: create_bin_dir
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/retry -trimpath \
	-ldflags "-s -w \
	-X main.BuildCommit=$(GIT_COMMIT) \
	-X main.BuildVersion=$(GIT_TAG) \
	-X main.BuildDate=$(BUILD_DATE) \
	-extldflags=-static" cmd/retry/main.go	

.PHONY: install
install: build
#	go install github.com/denkhaus/retry/cmd/retry@latest
	@cp bin/retry $(shell go env GOPATH)/bin
	@ls -la $(shell which retry)
	@retry -v