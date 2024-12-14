#!/bin/bash

# Function to display messages
function echo_msg {
  echo -e "\033[1;32m$1\033[0m"
}

# Function to prompt user
function prompt_user {
  read -p "$1 (y/n): " response
  if [[ "$response" =~ ^[Yy]$ ]]; then
    return 0  # Yes
  else
    return 1  # No
  fi
}

# Check if the script is running as root
if [[ "$EUID" -eq 0 ]]; then
  echo "Please do not run this script as root!"
  exit 1
fi

# Update macOS and install Xcode Command Line Tools
echo_msg "Updating macOS and installing Xcode Command Line Tools..."
xcode-select --install || echo "Xcode Command Line Tools already installed"

# Install Homebrew
if ! command -v brew &>/dev/null; then
  echo_msg "Installing Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
  echo_msg "Adding Homebrew to PATH..."
  SHELL_PROFILE="$HOME/.zprofile"
  touch "$SHELL_PROFILE"
  echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> "$SHELL_PROFILE"
  eval "$(/opt/homebrew/bin/brew shellenv)"
else
  echo_msg "Homebrew is already installed."
fi

# SSH Key Setup
if prompt_user "Set up a GitHub SSH key?"; then
  echo_msg "Setting up GitHub SSH key..."

  # Prompt user for GitHub email
  read -p "Enter your GitHub email address: " GITHUB_EMAIL

  # Generate SSH key
  ssh-keygen -t ed25519 -C "$GITHUB_EMAIL" -f "$HOME/.ssh/id_ed25519" -N "" && echo_msg "SSH key generated."

  # Start the ssh-agent
  eval "$(ssh-agent -s)"

  # Create or update SSH config file
  SSH_CONFIG="$HOME/.ssh/config"
  if [[ ! -f "$SSH_CONFIG" ]]; then
    touch "$SSH_CONFIG"
  fi

  # Add configuration for GitHub
  grep -qxF 'Host github.com' "$SSH_CONFIG" || cat >> "$SSH_CONFIG" <<EOL

Host github.com
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_ed25519
EOL

  # Add the key to the agent and macOS keychain
  ssh-add --apple-use-keychain "$HOME/.ssh/id_ed25519" && echo_msg "SSH key added to ssh-agent and keychain."

  # Copy public key to clipboard
  pbcopy < "$HOME/.ssh/id_ed25519.pub" && echo_msg "Public key copied to clipboard. Add it to your GitHub account."
  
  echo_msg "Go to https://github.com/settings/keys to add your new SSH key."
else
  echo_msg "Skipping SSH key setup."
fi

# Install Git
if prompt_user "Install Git?"; then
  echo_msg "Installing Git..."
  brew install git
else
  echo_msg "Skipping Git installation."
fi

# Install Visual Studio Code
if prompt_user "Install Visual Studio Code?"; then
  echo_msg "Installing Visual Studio Code..."
  brew install --cask visual-studio-code && echo_msg "Visual Studio Code installed successfully."
else
  echo_msg "Skipping Visual Studio Code installation."
fi

# Install Volta and Node.js
if prompt_user "Install Volta (for Node.js version management) and Node.js?"; then
  echo_msg "Installing Volta..."
  curl https://get.volta.sh | bash
  export VOLTA_HOME="$HOME/.volta"
  export PATH="$VOLTA_HOME/bin:$PATH"
  echo_msg "Installing Node.js (LTS version) using Volta..."
  volta install node
else
  echo_msg "Skipping Volta and Node.js installation."
fi

# Install Bun
if prompt_user "Install Bun?"; then
  echo_msg "Installing Bun..."
  curl -fsSL https://bun.sh/install | bash
else
  echo_msg "Skipping Bun installation."
fi

# Install Deno
if prompt_user "Install Deno?"; then
  echo_msg "Installing Deno..."
  brew install deno
else
  echo_msg "Skipping Deno installation."
fi

# Install pnpm
if prompt_user "Install pnpm?"; then
  echo_msg "Installing pnpm..."
  npm install -g pnpm
else
  echo_msg "Skipping pnpm installation."
fi

# Install and configure WezTerm
if prompt_user "Install WezTerm?"; then
  echo_msg "Installing WezTerm..."
  brew install --cask wezterm

  if prompt_user "Would you like to configure WezTerm?"; then
    echo_msg "Configuring WezTerm..."
    echo_msg "Installing MesloLGS Nerd Font Mono..."
    brew tap homebrew/cask-fonts
    brew install --cask font-meslo-lg-nerd-font
    cp ./wezterm_config.lua "$HOME/.wezterm.lua"
    echo_msg "WezTerm configuration installed to ~/.wezterm.lua."
  fi
fi

# Install and configure Zsh
if prompt_user "Configure Zsh?"; then
  echo_msg "Installing Zsh tools..."
  brew install powerlevel10k zsh-autosuggestions zsh-syntax-highlighting fzf eza zoxide thefuck
  cat ./zsh_config_snippet.zsh >> "$HOME/.zshrc"
  echo_msg "Zsh configuration updated. Please reload your terminal or source ~/.zshrc."
fi

# Prompt user to choose between Zen Browser and Arc Browser
echo_msg "Would you like to install a browser?"
echo "1) Zen Browser"
echo "2) Arc Browser"
echo "3) Skip browser installation"
read -p "Enter the number of your choice: " browser_choice

case $browser_choice in
  1)
    echo_msg "Installing Zen Browser..."
    brew install --cask zen-browser && echo_msg "Zen Browser installed successfully."
    ;;
  2)
    echo_msg "Installing Arc Browser..."
    brew install --cask arc && echo_msg "Arc Browser installed successfully."
    ;;
  3)
    echo_msg "Skipping browser installation."
    ;;
  *)
    echo_msg "Invalid choice. Skipping browser installation."
    ;;
esac

# Final Steps
echo_msg "Setup complete!"