REPO_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BINDIR := $(shell go env GOPATH)/bin
PATH := $(BINDIR):$(PATH):$(PWD)/bin
SHELL := env PATH=$(PATH) /bin/bash

PACKAGES := $(cd pkg && go list ./... | grep -v '/vendor/)

COVER_FILE := /tmp/telegraph-coverage.out
COVERAGE_OPTS := -coverprofile $(COVER_FILE)


all:	build

build:	dependencies codegen gobuild
	@echo "  - Build completed successfully."

check:	lint
lint:	protolint shellcheck golint govet
	@echo "  - Linting test scripts ..."
	@echo "  - Passed lint checks."

test:	tests
tests:	lint gotest
	@echo "  - Stay tuned for unit tests ..."

clean:
	@echo "  - Cleaning build ..."
	@echo "  - Cleaning protobuf build ..."
	(cd proto && $(MAKE) clean)


#
#  Handle dependencies - protoc, protolint.
#
deps:	dependencies

dependencies:
	@echo "  - Checking dependencies ..."
	(cd build && $(MAKE))

depclean:
	@echo "  - Cleaning dependencies ..."
	(cd build && $(MAKE) clean)


#
#  Generate and build golang code.
#
codegen:
	@echo "  - Generating golang code for protobuf ..."
	(cd proto && $(MAKE))

gobuild:
	@echo "  - Building golang code ..."
	@go mod download
	@go mod tidy


# Linters and tests.

#  Linters.
#
protolint:
	@echo "  - Linting protobuf ..."
	(cd proto && $(MAKE) lint)

shelllint:
	@echo "  - Linting shell scripts ..."
	(cd build && $(MAKE) lint)

golint:
	@echo "  - Linting golang code ..."
	gofmt -s -w .

govet:
	#@echo "  - Vetting golang code ..."
	#@(cd pkg && go vet ./...)

gotest:
	@echo "  - Testing golang code ..."
	@(cd pkg && go test $(COVERAGE_OPTS) ./...)

$(COVER_FILE):	gotest

coverage:	$(COVER_FILE)
	@echo "  - Generating coverage report ..."
	@(go tool cover -func=$(COVER_FILE) | sed 's#github\.com/.*/go-grpc-telegraph/#./#g')


.PHONY:	build check lint test tests clean
.PHONY:	deps dependencies depclean codegen gobuild
.PHONY:	protolint shelllint shellcheck golint govet gotest coverage
