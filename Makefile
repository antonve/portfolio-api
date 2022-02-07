.DEFAULT_GOAL := all
SHELL := /bin/bash -Eeuo pipefail

.PHONY: all
all: test run

.PHONY: test
test:
	@go test -v ./... -count=1

.PHONY: run
run:
	@go run cmd/server/main.go
