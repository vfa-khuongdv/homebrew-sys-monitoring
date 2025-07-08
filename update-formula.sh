#!/bin/bash

# Script to update Homebrew formula when releasing a new version
# Usage: ./update-formula.sh <version> <sha256>
# Example: ./update-formula.sh v1.0.0 abc123def456...

set -e

if [ $# -ne 2 ]; then
    echo "Usage: $0 <version> <sha256>"
    echo "Example: $0 v1.0.0 abc123def456..."
    exit 1
fi

VERSION=$1
SHA256=$2

# Remove 'v' prefix if present
VERSION_NUMBER=${VERSION#v}

FORMULA_FILE="Formula/sys-monitoring.rb"

# Update the formula file
sed -i.bak \
    -e "s|url \".*\"|url \"https://github.com/vfa-khuongdv/sys-monitoring/archive/refs/tags/${VERSION}.tar.gz\"|" \
    -e "s|sha256 \".*\"|sha256 \"${SHA256}\"|" \
    "$FORMULA_FILE"

# Remove backup file
rm "${FORMULA_FILE}.bak"

echo "Updated formula for version ${VERSION}"
echo "Don't forget to:"
echo "1. Test the formula: brew install --build-from-source ./Formula/sys-monitoring.rb"
echo "2. Commit and push the changes to your homebrew tap repository"
