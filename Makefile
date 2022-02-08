.DEFAULT_GOAL := all
SHELL := /bin/bash -Eeuo pipefail

.PHONY: all
all: test run

.PHONY: gen
gen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	oapi-codegen api/openapi.yaml > ports/openapi/openapi.gen.go

.PHONY: test
test:
	@go test -v ./... -count=1

.PHONY: run
run:
	@go run main.go
