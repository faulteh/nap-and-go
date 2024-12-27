# Variables
APP_NAME=nap-and-go
DOCKER_IMAGE=$(APP_NAME):latest
DOCKER_COMPOSE_FILE=docker-compose.yaml

# Go build parameters
GO_BUILD_CMD=go build
GO_FILES=./...
CGO_ENABLED=0

# Targets
.PHONY: all build bot web run lint test clean docker docker-compose up down info

# Default target
all: build

# Build both bot and web binaries
build: bot

# Build the Discord bot binary
bot:
	CGO_ENABLED=$(CGO_ENABLED) $(GO_BUILD_CMD) -o bin/bot ./cmd/bot

# Build the web interface binary
web:
	CGO_ENABLED=$(CGO_ENABLED) $(GO_BUILD_CMD) -o bin/web ./cmd/web

# Run both bot and web
run: build
	@echo "Starting bot and web interface..."
	./bin/bot & ./bin/web

# Lint the code
lint:
	golangci-lint run ./...

# Test the code
test:
	go test -v $(GO_FILES)

# Clean up binaries and other artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/

# Build Docker image
docker:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run using Docker Compose
docker-compose:
	@echo "Running Docker Compose..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

# Start Docker Compose
up:
	@echo "Starting services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Stop Docker Compose
down:
	@echo "Stopping services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down


# ------------------------------------------------------------------------------
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

VERSION_METADATA = unreleased

# Clear the "unreleased" string in BuildMetadata
ifneq ($(GIT_TAG),)
	VERSION_METADATA =
endif

info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"
	