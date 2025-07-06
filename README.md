# Quick Forge

Quickly set up a new MacBook with everything you need to develop in the JavaScript/TypeScript ecosystem.

## Installation

```bash
curl -fsSL https://raw.githubusercontent.com/tsotimus/quick-forge/main/install.sh | bash
```

## Usage

```bash
quickforge

# Dry run (see what would be installed)
quickforge -d
```

## Key Features

- One script to install and configure your development environment.
- Modular installs: choose what you need and skip what you don't.
- Pre-configured aliases for common git commands.

## What does the script do?

- Installs Homebrew, the macOS package manager.
- Installs Git and sets up an SSH key for GitHub.
- Installs Visual Studio Code (VSCode).
- Installs Node.js (via Fnm)
- Installs Bun (via Bum)
- Installs pnpm (via Corepack)
- Installs Warp, the AI Terminal.
- Lets you choose between installing Zen Browser or Arc Browser.

### Git Aliases

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

## Testing

QuickForge is designed specifically for macOS and uses Homebrew for installations. Testing approaches:

### GitHub Actions (Recommended)
```bash
# Trigger basic tests on macOS runners
gh workflow run test.yml

# Trigger comprehensive E2E tests (manual)
gh workflow run e2e-test.yml -f test_level=dry-run
gh workflow run e2e-test.yml -f test_level=safe
gh workflow run e2e-test.yml -f test_level=full  # Caution: installs software
```

### Local Testing
```bash
# Safe dry-run testing (no installations)
make test-dry-run

# Local testing (may install software)
make test-local

# Unit tests only
make test
```

### Docker Testing (Limited)
```bash
# Docker testing has limitations - see test.sh for details
make test-docker
```

**Note**: Docker testing shows expected failures for macOS-specific tools like Chrome since Docker containers run Linux.

#### Future Enhancements

- Add support for other development ecosystems (Python, Go, etc.).
- Integration with Docker or Kubernetes for containerized setups.
- Optional database installation (PostgreSQL, MongoDB, etc)
