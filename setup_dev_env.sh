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
  
  # Dynamically get the user's home directory and append the required path
  SHELL_PROFILE="$HOME/.zprofile"
  touch "$SHELL_PROFILE"  # Ensure the file exists
  
  # Add Homebrew shellenv to the shell profile
  echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> "$SHELL_PROFILE"
  
  # Apply the changes for the current session
  eval "$(/opt/homebrew/bin/brew shellenv)"
else
  echo_msg "Homebrew is already installed."
fi

# Install Git
if prompt_user "Install Git?"; then
  echo_msg "Installing Git..."
  brew install git
else
  echo_msg "Skipping Git installation."
fi

# Open VSCode download page in Safari
if prompt_user "Open Visual Studio Code download page in Safari?"; then
  echo_msg "Opening Visual Studio Code download page..."
  open -a Safari "https://code.visualstudio.com/download"
else
  echo_msg "Skipping Visual Studio Code download page."
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
if prompt_user "Install WezTerm (a GPU-accelerated terminal emulator)?"; then
  echo_msg "Installing WezTerm..."
  brew install --cask wezterm

  if prompt_user "Would you like to configure WezTerm?"; then
    echo_msg "Configuring WezTerm..."
    
    # Install MesloLGS Nerd Font
    echo_msg "Installing MesloLGS Nerd Font Mono..."
    brew tap homebrew/cask-fonts
    brew install --cask font-meslo-lg-nerd-font
    
    # Copy the configuration
    cp ./wezterm_config.lua "$HOME/.wezterm.lua"
    echo_msg "WezTerm configuration installed to ~/.wezterm.lua."
  fi
fi

# Install and configure Zsh
if prompt_user "Configure Zsh?"; then
  echo_msg "Installing Zsh tools..."
  
  # Install required tools
  brew install powerlevel10k zsh-autosuggestions zsh-syntax-highlighting fzf eza zoxide thefuck

  # Append the Zsh configuration snippet to ~/.zshrc
  cat ./zsh_config_snippet.zsh >> "$HOME/.zshrc"
  echo_msg "Zsh configuration updated. Please reload your terminal or source ~/.zshrc."
fi

# Final Steps
echo_msg "Setup complete!"