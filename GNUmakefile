SHELL = bash
GOTOOLS = \
	github.com/tools/godep
BUILDTIME=$(date -u -d "@${SOURCE_DATE_EPOCH:-$(date +%s)}" --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
GIT_IMPORT=github.com/continuul/random-names/command/version
GITCOMMIT=$(git rev-parse --short HEAD)
GIT_DIRTY?=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X $(GIT_IMPORT).BuildTime=$(BUILDTIME)

export GOLDFLAGS

.PHONY: all
all: bin

.PHONY: bin
bin: tools
	go build .

.PHONY: clean
clean:
	go clean .

.PHONY: tools
tools:
	go get -u -v $(GOTOOLS)
