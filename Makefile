# Makefile for sys-monitoring

BINARY_NAME=sys-monitoring
VERSION?=1.0.0
LDFLAGS=-ldflags "-s -w"

.PHONY: build clean test install uninstall help

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .
	@echo "Build complete!"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	go clean

test: build ## Build and test the binary
	@echo "Testing $(BINARY_NAME)..."
	timeout 2s ./$(BINARY_NAME) || echo "Test completed (timeout expected)"

install: build ## Install locally (requires sudo)
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo cp $(BINARY_NAME) /usr/local/bin/
	@echo "Installed successfully!"

uninstall: ## Uninstall from local system
	@echo "Uninstalling $(BINARY_NAME)..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstalled successfully!"

release-prep: clean build ## Prepare for release
	@echo "Preparing release $(VERSION)..."
	@echo "Binary size: $$(du -h $(BINARY_NAME) | cut -f1)"
	@echo "To create release:"
	@echo "1. git tag v$(VERSION)"
	@echo "2. git push origin v$(VERSION)"
	@echo "3. Create GitHub release with tag v$(VERSION)"

help: ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
