# Quick Forge

Quickly set up a new MacBook with everything you need to develop in the JavaScript/TypeScript ecosystem.

## Installation

### Option 1: One-line install (Recommended)
```bash
curl -fsSL https://raw.githubusercontent.com/tsotimus/quick-forge/main/install.sh | bash
```

### Option 2: Download binary manually
1. Go to the [releases page](https://github.com/tsotimus/quick-forge/releases)
2. Download the appropriate binary for your system
3. Make it executable: `chmod +x quickforge-darwin-arm64`
4. Move to your PATH: `sudo mv quickforge-darwin-arm64 /usr/local/bin/quickforge`

### Option 3: Build from source
```bash
git clone https://github.com/RockiRider/quick-forge.git
cd quickforge
make build
./quickforge
```

### Option 4: Go install (if you have Go installed)
```bash
go install github.com/RockiRider/quick-forge@latest
```

## Usage

```bash
# Interactive mode
quickforge

# Non-interactive mode (accepts all defaults)
quickforge -y

# Dry run (see what would be installed)
quickforge -d
```

## Key Features
- One script to install and configure your development environment.
- Modular installs: choose what you need and skip what you don't.
- Pre-configured zsh for enhanced productivity.

![Quick Forge Terminal Screenshot](./screenshot.png)

## What does the script do?

- Installs Homebrew, the macOS package manager.
- Installs Git and sets up an SSH key for GitHub.
- Installs Visual Studio Code (VSCode).
- Installs Node.js (via Fnm)
- Installs Bun (via Bum)
- Installs pnpm (via Corepack)
- Installs Warp, the AI Terminal.
- Lets you choose between installing Zen Browser or Arc Browser.


### Git Alias's
```shell
alias g='git'                      # Shortcut to replace 'git' with 'g'
alias gs='git status'              # Check current branch status
alias ga='git add'                 # Stage specific files
alias gaa='git add --all'          # Stage all changes (tracked and untracked)
alias gc='git commit'              # Commit staged changes
alias gap='git add --patch'        # Interactive staging of changes (hunks)
alias gp='git push'                # Push commits to the remote
alias gpl='git pull'               # Pull the latest changes from the remote
alias gl='git log'                 # Show commit history
alias gb='git branch'              # List or manage branches
alias gco='git checkout'           # Switch branches or restore files
alias gcon='git checkout -b'       # Checkout and create a new branch
alias gcm='git commit -m'          # Commit with a message inline
alias gundo='git reset --soft HEAD~1' # Undo the last commit (soft reset)
```

#### Future Enhancements

- Add support for other development ecosystems (Python, Go, etc.).
- Integration with Docker or Kubernetes for containerized setups.
- Optional database installation (PostgreSQL, MongoDB, etc)
