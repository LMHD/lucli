SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=lucli
BUILD_TIME=`date +%FT%T%z`

DEFAULT_SYSTEM_BINARY := $(BINARY).darwin

BINTRAY_API_KEY=$(shell cat api_key)
VERSION=$(shell cat VERSION)
BUILD_TIME=$(shell date +%FT%T%z)
BUILD_COMMIT=$(shell git rev-parse HEAD)

UNAME_S := $(shell uname -s)
DEFAULT_SHASUM_UTIL=shasum
ifeq ($(UNAME_S),Linux)
	DEFAULT_SHASUM_UTIL=sha1sum
endif

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags \"-X github.com/lmhd/lucli/lib.Version=${VERSION} -X github.com/lmhd/lucli/lib.BuildTime=${BUILD_TIME} -X github.com/lmhd/lucli/lib.BuildCommit=${BUILD_COMMIT}\"

DOCKER_RUN_COMMAND=docker run --rm -v $(shell pwd)/:/go/src/github.com/lmhd/lucli -w /go/src/github.com/lmhd/lucli

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(BINARY).darwin $(BINARY).linux.arm
	cp $(DEFAULT_SYSTEM_BINARY) $@

$(BINARY).darwin: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=darwin -e GOARCH=386 golang:1.9 /bin/bash -c "go get -v && go build ${LDFLAGS} -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha

$(BINARY).linux.arm: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=linux -e GOARCH=arm golang:1.9 /bin/bash -c "go get -v && go build ${LDFLAGS} -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha


.PHONY: clean
clean:
	rm -f -- ${BINARY}
	rm -f -- ${BINARY}.darwin
	rm -f -- ${BINARY}.linux.arm

.PHONY: install
install:
	cp ${BINARY}.darwin ~/bin/lucli

.PHONY: release
release:
	curl -T ${BINARY}.darwin -ulucymhdavies:${BINTRAY_API_KEY} https://api.bintray.com/content/lmhd/${BINARY}/${BINARY}/${VERSION}/${BINARY}-${VERSION}.darwin
	curl -T ${BINARY}.linux.arm -ulucymhdavies:${BINTRAY_API_KEY} https://api.bintray.com/content/lmhd/${BINARY}/${BINARY}/${VERSION}/${BINARY}-${VERSION}.linux.arm

