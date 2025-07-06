#!/bin/bash

# ⚠️  DOCKER LIMITATION NOTICE ⚠️
# This Docker-based test has limitations because QuickForge is designed for macOS
# and uses Homebrew to install macOS-specific applications like Google Chrome.
# Docker containers run Linux, which causes installations to fail with:
# "Error: macOS is required for this software."
#
# For proper testing, use GitHub Actions with macOS runners instead.
# See: .github/workflows/test.yml and .github/workflows/e2e-test.yml

echo "⚠️  WARNING: Docker testing has limitations for macOS-specific tools"
echo "🍎 QuickForge is designed for macOS and uses Homebrew for installations"
echo "🐧 Docker containers run Linux, causing macOS-specific installations to fail"
echo ""
echo "✅ For proper testing, use GitHub Actions with macOS runners:"
echo "   - .github/workflows/test.yml (basic testing)"
echo "   - .github/workflows/e2e-test.yml (comprehensive testing)"
echo ""
echo "🧪 Running Docker test anyway (will show limitations)..."
echo ""

docker run -it --rm quickforge zsh -i -c "
  echo '\n🔧 Step 1: Initial run (expect failures for macOS-specific tools)';
  /app/quickforge -y;

  echo '\n🔧 Step 2: Source shell and run again';
  source /root/.zshrc;
  /app/quickforge -y;

  echo '\n⚠️  E2E complete (with expected macOS compatibility issues)';
"

echo ""
echo "📝 Note: Chrome and other macOS-specific installations failed as expected"
echo "🚀 Use 'gh workflow run test.yml' for proper macOS testing"