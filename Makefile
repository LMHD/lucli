SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=lucli
BUILD_TIME=`date +%FT%T%z`

DEFAULT_SYSTEM_BINARY := $(BINARY).darwin.amd64

GO_VERSION := 1.13

BINTRAY_API_KEY=$(shell cat api_key)
VERSION=$(shell cat VERSION)
BUILD_TIME=$(shell date +%FT%T%z)
BUILD_COMMIT=$(shell git rev-parse HEAD)
BUILD_REPO=$(shell git remote get-url origin)
WHOAMI=$(shell whoami)

UNAME_S := $(shell uname -s)
DEFAULT_SHASUM_UTIL=shasum
ifeq ($(UNAME_S),Linux)
	DEFAULT_SHASUM_UTIL=sha1sum
	DEFAULT_SYSTEM_BINARY := $(BINARY).linux.amd64
endif

ifndef TRAVIS
	# TODO: consider enabling this
	DOCKER_RUN_COMMAND=docker run --rm -v $(shell pwd)/../../../:/go/src/ -w /go/src/github.com/lmhd/lucli
	#DOCKER_RUN_COMMAND=docker run --rm -v $(shell pwd)/:/go/src/github.com/lmhd/lucli -w /go/src/github.com/lmhd/lucli
	GITHUB_API_KEY=$(shell cat github_api)

	# Setup the -ldflags option for go build here, interpolate the variable values
	LDFLAGS=-ldflags \"-X github.com/lmhd/lucli/lib.Version=${VERSION} -X github.com/lmhd/lucli/lib.BuildTime=${BUILD_TIME} -X github.com/lmhd/lucli/lib.BuildCommit=${BUILD_COMMIT} -X github.com/lmhd/lucli/lib.BuildRepo=${BUILD_REPO} -X github.com/lmhd/lucli/lib.BuildUser=${WHOAMI}\"
endif
ifdef TRAVIS
	DOCKER_RUN_COMMAND=docker run --rm -v $(shell pwd)/:/go/src/github.com/lmhd/lucli -w /go/src/github.com/lmhd/lucli

	# Setup the -ldflags option for go build here, interpolate the variable values
	LDFLAGS=-ldflags \"-X github.com/lmhd/lucli/lib.Version=${VERSION} -X github.com/lmhd/lucli/lib.BuildTime=${BUILD_TIME} -X github.com/lmhd/lucli/lib.BuildCommit=${BUILD_COMMIT} -X github.com/lmhd/lucli/lib.BuildRepo=${BUILD_REPO} -X github.com/lmhd/lucli/lib.BuildUser=TravisCI\"
endif



.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(BINARY).darwin.amd64 $(BINARY).linux.amd64 $(BINARY).linux.arm
	cp $(DEFAULT_SYSTEM_BINARY) $@

$(BINARY).darwin.amd64: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=darwin -e GOARCH=amd64 golang:${GO_VERSION} /bin/bash -c "go get -v && go build ${LDFLAGS} -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha

$(BINARY).linux.amd64: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=linux -e GOARCH=amd64 golang:${GO_VERSION} /bin/bash -c "go get -v && go build ${LDFLAGS} -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha

$(BINARY).linux.arm: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=linux -e GOARCH=arm golang:${GO_VERSION} /bin/bash -c "go get -v && go build ${LDFLAGS} -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha


.PHONY: clean
clean:
	rm -f -- ${BINARY}
	rm -f -- ${BINARY}.darwin.amd64 ${BINARY}.darwin.amd64.sha
	rm -f -- ${BINARY}.linux.amd64  ${BINARY}.linux.amd64.sha
	rm -f -- ${BINARY}.linux.arm    ${BINARY}.linux.arm.sha

.PHONY: install
install:
	cp $(DEFAULT_SYSTEM_BINARY) ~/bin/lucli

.PHONY: release
release:
	./$(DEFAULT_SYSTEM_BINARY) github-release "${VERSION}" lucli.darwin.amd64 lucli.linux.amd64 lucli.linux.arm -- --github-access-token ${GITHUB_API_KEY} --github-repository lmhd/lucli

# Really simple "does it at least run?" tests for now
# Proper tests coming at some point
.PHONY: test
test: test-unit test-integration test-binary

.PHONY: test-unit
test-unit:
	echo "Coming soon"

.PHONY: test-integration
test-integration:
	go run main.go -d version --check-update=false
	go run main.go -d terraform version
	echo "More coming at some point"

.PHONY: test-binary
test-binary:
	./$(DEFAULT_SYSTEM_BINARY) -d version --check-update=false
	./$(DEFAULT_SYSTEM_BINARY) -d terraform version
