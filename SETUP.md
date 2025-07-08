# Homebrew Publishing Setup Guide

This guide will help you publish your `sys-monitoring` tool to Homebrew.

## Prerequisites

1. Your main repository: `https://github.com/vfa-khuongdv/sys-monitoring`
2. Your Homebrew tap repository: `https://github.com/vfa-khuongdv/homebrew-sys-monitoring`

## Step-by-Step Setup

### 1. Prepare Your Main Repository

First, make sure your main repository has the necessary files:

```bash
# In your sys-monitoring repository
cp Formula/sys-monitoring.rb /path/to/homebrew-sys-monitoring/Formula/
cp .github/workflows/update-homebrew.yml /path/to/sys-monitoring/.github/workflows/
```

### 2. Create a Release

1. Go to your main repository: `https://github.com/vfa-khuongdv/sys-monitoring`
2. Click "Releases" → "Create a new release"
3. Tag version: `v1.0.0`
4. Release title: `System Monitor v1.0.0`
5. Description: Describe your features
6. Click "Publish release"

### 3. Update the Formula with Correct SHA256

After creating the release, you need to get the SHA256 hash:

```bash
# Download the release tarball and calculate SHA256
curl -sL https://github.com/vfa-khuongdv/sys-monitoring/archive/refs/tags/v1.0.0.tar.gz | shasum -a 256

# Copy the hash and update your formula
```

### 4. Copy Files to Your Homebrew Tap Repository

Copy these files to your `homebrew-sys-monitoring` repository:

```bash
# File structure in homebrew-sys-monitoring:
homebrew-sys-monitoring/
├── Formula/
│   └── sys-monitoring.rb
├── README.md
└── LICENSE
```

### 5. Update the Formula

Edit `Formula/sys-monitoring.rb` and replace `PLACEHOLDER_SHA256` with the actual SHA256 hash from step 3.

### 6. Test Your Formula

```bash
# Test the formula locally
brew install --build-from-source ./Formula/sys-monitoring.rb

# Test that it works
sys-monitoring

# Uninstall for cleanup
brew uninstall sys-monitoring
```

### 7. Publish Your Tap

```bash
# In your homebrew-sys-monitoring repository
git add .
git commit -m "Add sys-monitoring formula v1.0.0"
git push origin main
```

## Usage for End Users

Once published, users can install your tool with:

```bash
# Add your tap
brew tap vfa-khuongdv/sys-monitoring

# Install the tool
brew install sys-monitoring

# Or in one command
brew install vfa-khuongdv/sys-monitoring/sys-monitoring
```

## Updating for New Releases

When you release a new version:

1. Create a new release in your main repository
2. The GitHub Action will automatically update the formula
3. Or manually use the `update-formula.sh` script:

```bash
./update-formula.sh v1.1.0 <new-sha256-hash>
```

## Troubleshooting

### Common Issues

1. **SHA256 mismatch**: Make sure you're using the correct hash from the release tarball
2. **Binary name mismatch**: Ensure your Go build creates a binary named `sys-monitoring`
3. **Formula syntax**: Use `brew audit Formula/sys-monitoring.rb` to check syntax

### Testing

```bash
# Audit your formula
brew audit --strict Formula/sys-monitoring.rb

# Test installation
brew install --build-from-source Formula/sys-monitoring.rb

# Test the binary
sys-monitoring
```

## Files Created

This setup created the following files:

- `Formula/sys-monitoring.rb` - The Homebrew formula
- `README.md` - Documentation for your tap
- `update-formula.sh` - Script to update formula for new releases  
- `build.sh` - Build script for local development
- `.github/workflows/update-homebrew.yml` - Auto-update workflow
- `SETUP.md` - This setup guide

Copy the appropriate files to your repositories and follow the steps above!
