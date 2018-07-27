SHELL = bash
GOTOOLS = \
	github.com/golang/lint/golint \
	github.com/tools/godep
BUILD_TIME="$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')"
VERSION=1.0.0
GIT_IMPORT=github.com/continuul/random-names/command
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DESCRIBE?=$(shell git describe --tags --always)
GIT_DIRTY?=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X $(GIT_IMPORT).GitDescribe=$(GIT_DESCRIBE) -X $(GIT_IMPORT).BuildTime=$(BUILD_TIME) -X $(GIT_IMPORT).Version=$(VERSION)
PACKAGES=$(shell go list ./... | grep -v '/vendor/')

export GIT_COMMIT
export GIT_DIRTY
export GIT_DESCRIBE
export GOLDFLAGS
export BUILD_TIME

CGO_ENABLED=0

.PHONY: all
all: bin

.PHONY: bin
bin:
	go build .

.PHONY: install
install: tools
	go install -ldflags "${GOLDFLAGS}" .
	GOOS=linux GOARCH=amd64 go install -ldflags "${GOLDFLAGS}" .

.PHONY: clean
clean:
	go clean .

.PHONY: ensure
ensure:
	dep ensure

.PHONY: format
format:
	go fmt $(PACKAGES)

.PHONY: lint
lint:
	@go list ./... \
		| grep -v /vendor/ \
		| cut -d '/' -f 4- \
		| xargs -n1 \
			golint ;\

.PHONY: tools
tools:
	go get -u -v $(GOTOOLS)

.PHONY: image
image:
	cp $(GOPATH)/bin/linux_amd64/random-names .
	docker build --build-arg VERSION=$(VERSION) -t continuul/names-generator:latest .
	rm random-names

