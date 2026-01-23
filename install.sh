#!/bin/bash
set -e

# Google Drive CLI Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/dl-alexandre/Google-Drive-CLI/master/install.sh | bash

VERSION="${GDRIVE_VERSION:-latest}"
INSTALL_DIR="${GDRIVE_INSTALL_DIR:-$HOME/.local/bin}"
REPO="dl-alexandre/Google-Drive-CLI"

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

case "$OS" in
    darwin)
        OS="darwin"
        ;;
    linux)
        OS="linux"
        ;;
    mingw*|msys*|cygwin*)
        OS="windows"
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

echo "Google Drive CLI Installer"
echo "=========================="
echo ""
echo "Detected: ${OS}/${ARCH}"
echo "Install directory: ${INSTALL_DIR}"
echo ""

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Check if Go is available for building from source
if command -v go &> /dev/null; then
    echo "Go detected. Building from source..."
    
    # Create temp directory
    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    
    # Clone and build
    echo "Cloning repository..."
    git clone --depth 1 "https://github.com/${REPO}.git" gdrive 2>/dev/null || {
        echo "Failed to clone repository. Trying local build..."
        cd -
        if [ -f "go.mod" ] && [ -f "cmd/gdrive/main.go" ]; then
            echo "Building from current directory..."
            go build -o "$INSTALL_DIR/gdrive" ./cmd/gdrive
        else
            echo "Error: Could not find source files"
            exit 1
        fi
    }
    
    if [ -d "$TEMP_DIR/gdrive" ]; then
        cd "$TEMP_DIR/gdrive"
        echo "Building..."
        go build -o "$INSTALL_DIR/gdrive" ./cmd/gdrive
        cd -
        rm -rf "$TEMP_DIR"
    fi
else
    echo "Go not found. Attempting to download pre-built binary..."
    
    # Construct download URL
    if [ "$VERSION" = "latest" ]; then
        DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/gdrive_${OS}_${ARCH}"
    else
        DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/gdrive_${OS}_${ARCH}"
    fi
    
    # Download binary
    echo "Downloading from: $DOWNLOAD_URL"
    if command -v curl &> /dev/null; then
        curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/gdrive" || {
            echo ""
            echo "Pre-built binary not available. Please install Go and run this script again."
            echo "Install Go from: https://go.dev/dl/"
            exit 1
        }
    elif command -v wget &> /dev/null; then
        wget -q "$DOWNLOAD_URL" -O "$INSTALL_DIR/gdrive" || {
            echo ""
            echo "Pre-built binary not available. Please install Go and run this script again."
            echo "Install Go from: https://go.dev/dl/"
            exit 1
        }
    else
        echo "Error: curl or wget required"
        exit 1
    fi
fi

# Make executable
chmod +x "$INSTALL_DIR/gdrive"

# Verify installation
if [ -x "$INSTALL_DIR/gdrive" ]; then
    echo ""
    echo "Installation successful!"
    echo ""
    "$INSTALL_DIR/gdrive" version 2>/dev/null || echo "gdrive installed to $INSTALL_DIR/gdrive"
    echo ""
    
    # Check if install dir is in PATH
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo "Add the following to your shell profile (.bashrc, .zshrc, etc.):"
        echo ""
        echo "  export PATH=\"\$PATH:$INSTALL_DIR\""
        echo ""
    fi
    
    echo "Quick start:"
    echo "  gdrive auth login    # Authenticate with Google Drive"
    echo "  gdrive files list    # List your files"
    echo "  gdrive --help        # See all commands"
else
    echo "Installation failed"
    exit 1
fi
