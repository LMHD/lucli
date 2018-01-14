SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=lucli
BUILD_TIME=`date +%FT%T%z`

DEFAULT_SYSTEM_BINARY := $(BINARY).darwin-amd64

BINTRAY_API_KEY=$(shell cat api_key)
VERSION=$(shell cat VERSION)

UNAME_S := $(shell uname -s)
DEFAULT_SHASUM_UTIL=shasum
ifeq ($(UNAME_S),Linux)
	DEFAULT_SHASUM_UTIL=sha1sum
endif

DOCKER_RUN_COMMAND=docker run --rm -v $(shell pwd)/:/go/src/github.com/lmhd/lucli -w /go/src/github.com/lmhd/lucli

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(BINARY).darwin-amd64 $(BINARY).linux-arm
	cp $(DEFAULT_SYSTEM_BINARY) $@

$(BINARY).darwin-amd64: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=darwin-amd64 golang:1.9 /bin/bash -c "go get -v && go build -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha

$(BINARY).linux-arm: $(SOURCES)
	${DOCKER_RUN_COMMAND} -e GOOS=linux -e GOARCH=arm golang:1.9 /bin/bash -c "go get -v && go build -o $@"
	${DEFAULT_SHASUM_UTIL} $@ > $@.sha


.PHONY: clean
clean:
	rm -f -- ${BINARY}
	rm -f -- ${BINARY}.darwin-amd64
	rm -f -- ${BINARY}.linux-arm

.PHONY: install
install:
	cp ${BINARY}.darwin-amd64 ~/bin/lucli

.PHONY: release
release:
	curl -T ${BINARY}.darwin-amd64 -ulucymhdavies:${BINTRAY_API_KEY} https://api.bintray.com/content/lmhd/${BINARY}/${BINARY}/${VERSION}/${BINARY}.darwin-amd64
	curl -T ${BINARY}.linux-arm    -ulucymhdavies:${BINTRAY_API_KEY} https://api.bintray.com/content/lmhd/${BINARY}/${BINARY}/${VERSION}/${BINARY}.linux-arm

