#!/bin/bash

# Function to display messages
function echo_msg {
  echo -e "\033[1;31m$1\033[0m"
}

# Check if the script is running as root
if [[ "$EUID" -eq 0 ]]; then
  echo "Please do not run this script as root!"
  exit 1
fi

# Uninstall Docker
if brew list --cask | grep -q "^docker\$"; then
  echo_msg "Uninstalling Docker..."
  brew uninstall --cask docker
else
  echo_msg "Docker is not installed."
fi

# Uninstall pnpm
if command -v pnpm &>/dev/null; then
  echo_msg "Uninstalling pnpm..."
  npm uninstall -g pnpm
else
  echo_msg "pnpm is not installed."
fi

# Uninstall Bun
if command -v bun &>/dev/null; then
  echo_msg "Uninstalling Bun..."
  rm -rf ~/.bun
else
  echo_msg "Bun is not installed."
fi

# Uninstall Deno
if brew list | grep -q "^deno\$"; then
  echo_msg "Uninstalling Deno..."
  brew uninstall deno
else
  echo_msg "Deno is not installed."
fi

# Uninstall Node.js and nvm
if [ -d "$HOME/.nvm" ]; then
  echo_msg "Uninstalling nvm and Node.js..."
  rm -rf "$HOME/.nvm"
  sed -i '' '/export NVM_DIR/d' ~/.bashrc ~/.zshrc
  sed -i '' '/\[ -s "\$NVM_DIR\/nvm.sh" \]/d' ~/.bashrc ~/.zshrc
else
  echo_msg "nvm is not installed."
fi

# Uninstall VSCode
if brew list --cask | grep -q "^visual-studio-code\$"; then
  echo_msg "Uninstalling Visual Studio Code..."
  brew uninstall --cask visual-studio-code
else
  echo_msg "Visual Studio Code is not installed."
fi

# Uninstall Git
if brew list | grep -q "^git\$"; then
  echo_msg "Uninstalling Git..."
  brew uninstall git
else
  echo_msg "Git is not installed."
fi

# Uninstall Homebrew
if command -v brew &>/dev/null; then
  echo_msg "Uninstalling Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/uninstall.sh)"
else
  echo_msg "Homebrew is not installed."
fi

# Cleanup
echo_msg "Removing leftover configuration files..."
rm -rf ~/.bashrc ~/.zshrc ~/.npm ~/.pnpm ~/.config/pnpm

# Final Message
echo_msg "Uninstallation complete. Your system is reverted to its original state."