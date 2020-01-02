.PHONY: help lint test race cover coverhtml

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "      lint               to run golint on files recursively"
	@echo "      test               to run tests"
	@echo "      race               to run tests with race detector"
	@echo "      cover              to run tests with coverage"
	@echo "      coverhtml          to run tests with coverage and generate html output"

lint:
	@set -e; \
	golangci-lint -v run

test:
	@go test -v ./...

race:
	@go test -race -v ./...

cover:
	@go test -v -coverprofile=coverage.out -cover ./...

coverhtml:
	@go test -v -coverprofile=coverage.out -cover ./...
	@go tool cover -html=coverage.out
