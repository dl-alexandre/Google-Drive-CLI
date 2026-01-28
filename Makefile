VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# Load bundled OAuth credentials from .env if present (for development builds)
-include .env
export

GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOMOD = $(GOCMD) mod

BINARY_NAME = gdrv
BINARY_DIR = bin

# Inject OAuth credentials if available (from .env or environment)
OAUTH_LDFLAGS =
ifdef GDRV_CLIENT_ID
	OAUTH_LDFLAGS += -X github.com/dl-alexandre/gdrv/internal/auth.BundledOAuthClientID=$(GDRV_CLIENT_ID)
endif
ifdef GDRV_CLIENT_SECRET
	OAUTH_LDFLAGS += -X github.com/dl-alexandre/gdrv/internal/auth.BundledOAuthClientSecret=$(GDRV_CLIENT_SECRET)
endif

LDFLAGS = -ldflags "-X github.com/dl-alexandre/gdrv/pkg/version.Version=$(VERSION) \
	-X github.com/dl-alexandre/gdrv/pkg/version.GitCommit=$(GIT_COMMIT) \
	-X github.com/dl-alexandre/gdrv/pkg/version.BuildTime=$(BUILD_TIME) \
	$(OAUTH_LDFLAGS)"

PLATFORMS = linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

.PHONY: all build clean test deps tidy lint install help

all: deps build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) ./cmd/gdrv

build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BINARY_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*} GOARCH=$${platform#*/} \
		$(GOBUILD) $(LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}$(if $(findstring windows,$${platform}),.exe,) ./cmd/gdrv; \
		echo "Built $(BINARY_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}"; \
	done

deps:
	@echo "Installing dependencies..."
	$(GOMOD) download

tidy:
	@echo "Tidying go modules..."
	$(GOMOD) tidy

test:
	@echo "Running tests..."
	$(GOTEST) -v -race -cover ./...

test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)
	rm -f coverage.out coverage.html

install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BINARY_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

checksums:
	@echo "Generating checksums..."
	@cd $(BINARY_DIR) && \
	for file in $(BINARY_NAME)*; do \
		if [ -f "$$file" ]; then \
			shasum -a 256 "$$file" >> checksums.txt; \
		fi \
	done
	@echo "Checksums written to $(BINARY_DIR)/checksums.txt"

version:
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"

run:
	@$(GOBUILD) $(LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) ./cmd/gdrv
	@./$(BINARY_DIR)/$(BINARY_NAME) $(ARGS)

help:
	@echo "Google Drive CLI (gdrv) Makefile"
	@echo ""
	@echo "Targets:"
	@echo "  all          - Build the binary (default)"
	@echo "  build        - Build for current platform"
	@echo "  build-all    - Build for all platforms"
	@echo "  deps         - Download dependencies"
	@echo "  tidy         - Tidy go modules"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  lint         - Run linter"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install binary to GOPATH/bin"
	@echo "  checksums    - Generate SHA256 checksums"
	@echo "  version      - Show version info"
	@echo "  run          - Build and run (use ARGS=... for arguments)"
	@echo "  help         - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make test"
	@echo "  make run ARGS='version'"
	@echo "  make build-all"
