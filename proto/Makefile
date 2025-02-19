#!/usr/bin/env make

PROTO_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BINDIR := $(shell go env GOPATH)/bin
PATH := $(BINDIR):$(PATH):$(PWD)/bin
SHELL := env PATH=$(PATH) /bin/bash

all:	build test

build:	dependencies codegen
	@echo "  - Protobuf build completed successfully."

check:	lint
lint:	protolint
	@echo "  - Passed lint checks."

test:	tests
tests:	lint
	@echo "  - Passed tests."

clean:
	@echo "  - Cleaning build ..."
	@$(RM) *.pb.go
	

#
#  Handle dependencies - protoc, protolint.
#
deps:	dependencies

dependencies:
	@echo "  - Checking dependencies ..."
	(cd ../build/dependencies && $(MAKE) install)

depclean:
	@echo "  - Cleaning dependencies ..."
	(cd ../build/dependencies && $(MAKE) clean)


#
#  Generate and build golang code.
#
codegen:	protoc
	@echo "  - Generating golang code for protobuf ..."
	@protoc --go_out=./ --go_opt=paths=source_relative            \
	        --go-grpc_out=./ --go-grpc_opt=paths=source_relative  \
	        -I "../build/dependencies/protobuf/include"  -I ./    \
	        *.proto

protoc:
	@echo "  - Checking protoc ..."
	(cd ../build/dependencies && $(MAKE) protoc)


# Linters and tests.

#  Linters.
#
protolint:	$(BINDIR)/protolint
	@echo "  - Linting protobuf ..."
	@echo "  - protolint version = $$(protolint version)"
	protolint lint *.proto

$(BINDIR)/protolint:
	@echo "  - Missing protolint, installing dependencies ..."
	@$(MAKE) dependencies


.PHONY:	build check lint test tests clean
.PHONY:	deps dependencies depclean codegen protoc protolint
