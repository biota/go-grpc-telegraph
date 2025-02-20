#!/usr/bin/env make

BUILD_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BINDIR := $(shell go env GOPATH)/bin
PATH := $(BINDIR):$(PATH):$(PWD)/bin
SHELL := env PATH=$(PATH) /bin/bash

# Use full path to utilities as /usr/bin/protoc could exist and on Ubuntu
# systems, it is a much older version.
UTILITIES := $(BINDIR)/protoc $(BINDIR)/protolint

include bin/releases.env
export


all:	install

build:	install
install:	deps
	@echo "  - Dependencies installed: $(UTILITIES)"

check:	tests
test:	tests
tests:	deps lint

lint:
	@echo "  - Linting shell scripts ..."
	shellcheck "$(BUILD_DIR)/bin/install-deps"

clean:
	@echo "  - Removing utilities ..."
	$(RM) -rf $(BINDIR)/protoc* $(BINDIR)/protolint

	@echo "  - Removing downloaded dependencies ..."
	$(RM) -rf "$(BUILD_DIR)"/dependencies/*

	@echo "  - Cleaned up build environment."


#
#  Handle dependencies - protolint, protoc.
#
deps:	$(UTILITIES)
	@echo "  - Checking if protoc version is the expected version ..."
	# grep-ping ensures if it fails, this target fails.
	@protoc --version | grep "protoc $(PROTOC_RELEASE)"
	@echo "  - protoc version = $$(protoc --version)"
	@echo "  - protolint version = $$(protolint version)"


$(UTILITIES):
	@echo "  - Checking if $@ dependency is installed ..."
	@command -v "$@" || ("$(BUILD_DIR)/bin/install-deps")


.PHONY:	build install check test tests lint clean deps
