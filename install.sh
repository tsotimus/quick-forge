#!/bin/bash
# QuickForge Installation Script
# Usage: curl -fsSL https://raw.githubusercontent.com/tsotimus/quickforge/main/install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Check if running on macOS
if [[ "$(uname -s)" != "Darwin" ]]; then
    print_error "This tool is designed for macOS only"
    exit 1
fi

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64) ARCH="arm64" ;;
    *) 
        print_error "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

print_status "Detected macOS with $ARCH architecture"

# Set variables
REPO="tsotimus/quickforge"
BINARY_NAME="quickforge-darwin-$ARCH"
INSTALL_DIR="/usr/local/bin"
INSTALL_PATH="$INSTALL_DIR/quickforge"

# Check if quickforge is already installed
if command -v quickforge &> /dev/null; then
    CURRENT_VERSION=$(quickforge --version 2>/dev/null || echo "unknown")
    print_warning "QuickForge is already installed (version: $CURRENT_VERSION)"
    read -p "Do you want to update it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_status "Installation cancelled"
        exit 0
    fi
fi

# Create temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

print_status "Downloading QuickForge..."

# Get the latest release URL
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$BINARY_NAME"

# Download the binary
if ! curl -fsSL "$DOWNLOAD_URL" -o quickforge; then
    print_error "Failed to download QuickForge from $DOWNLOAD_URL"
    print_error "Please check if the release exists and try again"
    exit 1
fi

# Make it executable
chmod +x quickforge

print_status "Installing QuickForge to $INSTALL_PATH..."

# Check if we need sudo
if [[ ! -w "$INSTALL_DIR" ]]; then
    print_status "Administrator privileges required to install to $INSTALL_DIR"
    if ! sudo mv quickforge "$INSTALL_PATH"; then
        print_error "Failed to install QuickForge"
        exit 1
    fi
else
    if ! mv quickforge "$INSTALL_PATH"; then
        print_error "Failed to install QuickForge"
        exit 1
    fi
fi

# Clean up
cd /
rm -rf "$TEMP_DIR"

# Verify installation
if command -v quickforge &> /dev/null; then
    print_success "QuickForge installed successfully!"
    echo
    echo "⚡ Run 'quickforge' to get started"
    echo "⚡ Run 'quickforge --help' for usage information"
    echo
    echo "🔨 QuickForge will help you set up your macOS development environment"
else
    print_error "Installation completed but quickforge command not found"
    print_error "You may need to add $INSTALL_DIR to your PATH"
    exit 1
fi 