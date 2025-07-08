#!/bin/bash

# Build script for sys-monitoring
# This ensures the binary is named correctly for Homebrew

set -e

# Build for current platform
echo "Building sys-monitoring..."
go build -ldflags "-s -w" -o sys-monitoring .

echo "Build complete!"
echo "Binary created: sys-monitoring"

# Test that it runs
if [[ "$1" == "--test" ]]; then
    echo "Testing binary..."
    timeout 2s ./sys-monitoring || echo "Test completed (timeout expected)"
fi
