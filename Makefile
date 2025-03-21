.PHONY: all build test release image

# Default target
all: deps build

# Install dependencies
deps:
	@gum log --level info "Installing dependencies..."
	@go mod tidy
	@go mod download

# Build all main packages
build:
	@gum log --level info "Building binaries..."
	@for dir in $$(rg -t go "package main" --files-with-matches | xargs -I{} dirname {} | sort -u); do \
		gum log --level info "Building $$dir"; \
		go build -o bin/$$(basename $$dir) ./$$dir; \
	done

# Run tests with gum logging
test:
	@gum log --level info "Running tests..."
	@gum log --level info "Starting tests"
	@go test -v ./... || (gum log --level error "Tests failed" && exit 1)
	@gum log --level debug "Tests completed successfully"

# Build Docker image if Dockerfile exists
image:
	@if [ -f Dockerfile ]; then \
		gum log --level info "Building Docker image..."; \
		docker build -t $$(basename $$(pwd)) .; \
	else \
		gum log --level info "Error: Dockerfile not found"; \
		exit 1; \
	fi

# Release using goreleaser
release:
	@gum log --level info "Preparing release..."
	@curl -s -o .goreleaser.yml https://raw.githubusercontent.com/goreleaser/goreleaser/main/www/docs/static/examples/go-binary.yaml
	@go install github.com/goreleaser/goreleaser@latest
	@goreleaser release --rm-dist
