include .makeup/makeup-kit-info/main.mk
include .makeup/makeup-kit-go-cli/main.mk

# OVERRIDES
GO_CLI_SRC_PACKAGES := . ./cmd
GO_CLI_ORG_NAME := deis

build: go-cli-build

test: go-cli-test

clean: go-cli-clean

lint: go-cli-lint

vet: go-cli-vet

fmt: go-cli-fmt

test-style: fmt lint vet

install: go-cli-install
