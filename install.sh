#!/bin/bash
# QuickForge Installation Script
# Usage: curl -fsSL https://raw.githubusercontent.com/tsotimus/quick-forge/main/install.sh | bash
# Or with auto-update: curl -fsSL https://raw.githubusercontent.com/tsotimus/quick-forge/main/install.sh | QUICKFORGE_UPDATE=yes bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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

REPO="tsotimus/quick-forge"
BINARY_NAME="quickforge-darwin-$ARCH"
INSTALL_DIR="/opt/quickforge/bin"
INSTALL_PATH="$INSTALL_DIR/quickforge"
PATH_MARKER="# Added by QuickForge"
PATH_LINE="export PATH=\"$INSTALL_DIR:\$PATH\""

get_shell_profile() {
    case "${SHELL##*/}" in
        zsh) echo "$HOME/.zprofile" ;;
        bash) echo "$HOME/.bash_profile" ;;
        *) echo "$HOME/.profile" ;;
    esac
}

ensure_install_dir() {
    if [[ -d "$INSTALL_DIR" ]]; then
        return 0
    fi

    print_status "Creating $INSTALL_DIR..."
    if ! sudo mkdir -p "$INSTALL_DIR"; then
        print_error "Failed to create $INSTALL_DIR"
        exit 1
    fi

    if ! sudo chown -R "$(whoami)" /opt/quickforge; then
        print_error "Failed to set ownership on /opt/quickforge"
        exit 1
    fi
}

install_binary() {
    local source_path="$1"

    if [[ ! -f "$source_path" ]]; then
        print_error "Downloaded binary not found at $source_path"
        exit 1
    fi

    print_status "Installing QuickForge to $INSTALL_PATH..."
    ensure_install_dir

    if [[ -w "$INSTALL_DIR" ]]; then
        if ! mv "$source_path" "$INSTALL_PATH"; then
            print_error "Failed to install QuickForge to $INSTALL_PATH"
            exit 1
        fi
    else
        print_status "Administrator privileges required to install to $INSTALL_DIR"
        if ! sudo mv "$source_path" "$INSTALL_PATH"; then
            print_error "Failed to install QuickForge to $INSTALL_PATH"
            exit 1
        fi
        sudo chown "$(whoami)" "$INSTALL_PATH"
    fi

    chmod +x "$INSTALL_PATH"
}

configure_path() {
    local profile
    profile="$(get_shell_profile)"

    if [[ -f "$profile" ]] && grep -qF "$INSTALL_DIR" "$profile" 2>/dev/null; then
        print_status "PATH already configured in ~/${profile##*/}"
        return 0
    fi

    print_status "Adding $INSTALL_DIR to PATH in ~/${profile##*/}..."
    {
        echo ""
        echo "$PATH_MARKER"
        echo "$PATH_LINE"
    } >> "$profile"
}

is_quickforge_installed() {
    [[ -x "$INSTALL_PATH" ]] || command -v quickforge &> /dev/null
}

# Check if quickforge is already installed
if is_quickforge_installed; then
    CURRENT_VERSION=$("$INSTALL_PATH" --version 2>/dev/null || quickforge --version 2>/dev/null || echo "unknown")
    print_warning "QuickForge is already installed (version: $CURRENT_VERSION)"

    if [[ "${QUICKFORGE_UPDATE}" == "yes" || "${QUICKFORGE_UPDATE}" == "y" ]]; then
        print_status "Auto-updating due to QUICKFORGE_UPDATE environment variable"
    elif [[ -t 0 ]]; then
        read -p "Do you want to update it? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            print_status "Installation cancelled"
            exit 0
        fi
    else
        print_status "Updating QuickForge (piped installation detected)"
        print_status "To skip this prompt in future, use: QUICKFORGE_UPDATE=yes"
    fi
fi

TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT
cd "$TEMP_DIR"

print_status "Downloading QuickForge..."

DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/$BINARY_NAME"

if ! curl -fsSL "$DOWNLOAD_URL" -o quickforge; then
    print_error "Failed to download QuickForge from $DOWNLOAD_URL"
    print_error "Please check if the release exists and try again"
    exit 1
fi

chmod +x quickforge
install_binary "$TEMP_DIR/quickforge"
configure_path

export PATH="$INSTALL_DIR:$PATH"

if [[ ! -x "$INSTALL_PATH" ]]; then
    print_error "Installation failed: binary missing at $INSTALL_PATH"
    exit 1
fi

print_success "QuickForge installed successfully!"
echo
echo "⚡ Run 'quickforge' to get started"
echo "⚡ Run 'quickforge --help' for usage information"
echo
echo "Installed to: $INSTALL_PATH"

if ! command -v quickforge &> /dev/null; then
    print_warning "Open a new terminal, or run:"
    echo "  source $(get_shell_profile)"
fi

echo
echo "🔨 QuickForge will help you set up your macOS development environment"
