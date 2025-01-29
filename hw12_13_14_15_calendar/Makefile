BIN := "./bin/calendar"

GIT_HASH := $(shell git log --format="%h" -n 1)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%S)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(BUILD_DATE) -X main.gitHash=$(GIT_HASH)
LINTER := golangci-lint
LINTER-CONFIG := ../.golangci.yml
GOBIN ?= $(shell go env GOPATH)/bin

.DEFAULT_GOAL := build

.PHONY: all lint install-linter clean build

all: lint
# Build the binary
build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

# Build and run the service with default config (optional)
run: build
	$(BIN) -config ./configs/calendar-config.yaml


lint: install-linter build
	@echo "Запуск golangci-lint"
	golangci-lint run --concurrency 1 ./...


install-linter:
	(which golangci-lint > /dev/null) || curl -sSfl https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1


# Run tests with race detector
test:
	go test -v -race ./internal/storage/...

# Apply database migrations (optional)
migrate:
	./migrations/apply.sh

# Clean up build artifacts
clean:
	rm -rf $(BIN)
