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
COVERAGE_DIR = .artifacts/coverage

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
	-X github.com/dl-alexandre/cli-tools/version.Version=$(VERSION) \
	-X github.com/dl-alexandre/cli-tools/version.GitCommit=$(GIT_COMMIT) \
	-X github.com/dl-alexandre/cli-tools/version.BuildTime=$(BUILD_TIME) \
	$(OAUTH_LDFLAGS)"

# Optimized build flags for smaller binaries
# -s: disable symbol table
# -w: disable DWARF debug info  
# These reduce binary size by ~30-35%
OPTIMIZED_LDFLAGS = -ldflags "-s -w -X github.com/dl-alexandre/gdrv/pkg/version.Version=$(VERSION) \
	-X github.com/dl-alexandre/gdrv/pkg/version.GitCommit=$(GIT_COMMIT) \
	-X github.com/dl-alexandre/gdrv/pkg/version.BuildTime=$(BUILD_TIME) \
	-X github.com/dl-alexandre/cli-tools/version.Version=$(VERSION) \
	-X github.com/dl-alexandre/cli-tools/version.GitCommit=$(GIT_COMMIT) \
	-X github.com/dl-alexandre/cli-tools/version.BuildTime=$(BUILD_TIME) \
	$(OAUTH_LDFLAGS)"

PLATFORMS = linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

.PHONY: all build clean test deps tidy lint security checksums version install help format install-hooks check vet

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

# Build optimized (smaller) binary - strips debug info
# Reduces size by ~30-35% (40MB → 27MB)
build-optimized:
	@echo "Building optimized $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	$(GOBUILD) -trimpath $(OPTIMIZED_LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) ./cmd/gdrv
	@echo "Built optimized binary: $(BINARY_DIR)/$(BINARY_NAME)"
	@echo "Size: $$(ls -lh $(BINARY_DIR)/$(BINARY_NAME) | awk '{print $$5}')"

# Build optimized for all platforms
build-all-optimized:
	@echo "Building optimized binaries for all platforms..."
	@mkdir -p $(BINARY_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*} GOARCH=$${platform#*/} \
		$(GOBUILD) -trimpath $(OPTIMIZED_LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}$(if $(findstring windows,$${platform}),.exe,) ./cmd/gdrv; \
		echo "Built optimized $(BINARY_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/} $$(ls -lh $(BINARY_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}$(if $(findstring windows,$${platform}),.exe,) 2>/dev/null | awk '{print $$5}')"; \
	done

# Compare binary sizes
size-compare:
	@echo "Building both versions for comparison..."
	@mkdir -p $(BINARY_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-debug ./cmd/gdrv
	$(GOBUILD) -trimpath $(OPTIMIZED_LDFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)-optimized ./cmd/gdrv
	@echo ""
	@echo "Binary Size Comparison:"
	@echo "  Debug:     $$(ls -lh $(BINARY_DIR)/$(BINARY_NAME)-debug | awk '{print $$5}')"
	@echo "  Optimized: $$(ls -lh $(BINARY_DIR)/$(BINARY_NAME)-optimized | awk '{print $$5}')"
	@echo "  Gzipped:   $$(gzip -c $(BINARY_DIR)/$(BINARY_NAME)-optimized | wc -c | awk '{printf "%.1f MB", $$1/1024/1024}')"
	@rm -f $(BINARY_DIR)/$(BINARY_NAME)-debug $(BINARY_DIR)/$(BINARY_NAME)-optimized

deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) verify

tidy:
	@echo "Tidying go modules..."
	$(GOMOD) tidy

# Run all checks (format, vet, lint, test)
.PHONY: check
check: format vet lint test

# Run go vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

test:
	@echo "Running tests..."
	$(GOTEST) -v -race -cover ./...

test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) -v -race -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	$(GOCMD) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report: $(COVERAGE_DIR)/coverage.html"

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
	rm -rf $(COVERAGE_DIR)

security:
	@echo "Running govulncheck..."
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck not installed. Run: go install golang.org/x/vuln/cmd/govulncheck@latest"; \
	fi

install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BINARY_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

checksums:
	@echo "Generating checksums..."
	@cd $(BINARY_DIR) && rm -f checksums.txt && \
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

format:
	@echo "Formatting code..."
	@gofmt -w -s .
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	else \
		echo "goimports not installed. Install: go install golang.org/x/tools/cmd/goimports@latest"; \
	fi

install-hooks:
	@echo "Installing git hooks..."
	@git config core.hooksPath .githooks
	@echo "Hooks installed from .githooks/"

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
