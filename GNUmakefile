SHELL = bash
GOTOOLS = \
	github.com/tools/godep
BUILDTIME="$(shell date -u)"
GIT_IMPORT=github.com/continuul/random-names/command
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DESCRIBE?=$(shell git describe --tags --always)
GIT_DIRTY?=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X $(GIT_IMPORT).GitDescribe=$(GIT_DESCRIBE)

export GIT_COMMIT
export GIT_DIRTY
export GIT_DESCRIBE
export GOLDFLAGS
export BUILDTIME

CGO_ENABLED=0

.PHONY: all
all: bin

.PHONY: bin
bin: tools
	go build .

.PHONY: install
install: tools
	echo $(BUILDTIME)
	go install -ldflags "${GOLDFLAGS}" .
	GOOS=linux GOARCH=amd64 go install -ldflags "${GOLDFLAGS}" .

.PHONY: clean
clean:
	go clean .

.PHONY: tools
tools:
	go get -u -v $(GOTOOLS)

.PHONY: image
image:
	cp $(GOPATH)/bin/linux_amd64/random-names .
	docker build -t continuul/names-generator:latest .
	rm random-names

