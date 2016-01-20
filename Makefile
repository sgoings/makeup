# makeup-managed:begin
include makeup.mk
# makeup-managed:end

include .makeup/makeup-bag-deis/info.mk
include .makeup/makeup-bag-deis/go-cli.mk

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

update: go-cli-glide-update
