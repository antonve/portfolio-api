.DEFAULT_GOAL := all
SHELL := /bin/bash -Eeuo pipefail

.PHONY: all
all: migrate test run

.PHONY: gen
gen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	oapi-codegen api/openapi.yaml > ports/openapi/openapi.gen.go

.PHONY: test
test:
	go test -v ./... -count=1

.PHONY: migrate
migrate:
	@SERVER_TO_RUN=migrate go run main.go

.PHONY: run
run:
	@SERVER_TO_RUN=http go run main.go
