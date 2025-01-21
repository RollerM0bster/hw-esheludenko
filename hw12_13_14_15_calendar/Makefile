BIN := "./bin/calendar"
DOCKER_IMG := "calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%S)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(BUILD_DATE) -X main.gitHash=$(GIT_HASH)

.DEFAULT_GOAL := build

# Build the binary
build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

# Build and run the service with default config (optional)
run: build
	$(BIN) -config ./configs/calendar-config.yaml

# Run tests with race detector
test:
	go test -v -race ./internal/storage/...


# Run golangci-lint
lint:
	golangci-lint run ./...

# Build Docker image
build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

# Run Docker image
run-img: build-img
	docker run $(DOCKER_IMG)

# Show the version of the built binary
version: build
	$(BIN) version

# Apply database migrations (optional)
migrate:
	./migrations/apply.sh

# Clean up build artifacts
clean:
	rm -rf $(BIN)

.PHONY: build run test lint build-img run-img version migrate clean install-lint-deps
