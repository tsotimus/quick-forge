package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

const GitAliases = `
# Git aliases
alias g='git'                      # Shortcut to replace 'git' with 'g'
alias gs='git status'              # Check current branch status
alias ga='git add'                 # Stage specific files
alias gaa='git add --all'         # Stage all changes (tracked and untracked)
alias gc='git commit'              # Commit staged changes
alias gap='git add --patch'       # Interactive staging of changes (hunks)
alias gp='git push'               # Push commits to the remote
alias gl='git log'                # Show commit history
alias gb='git branch'             # List or manage branches
alias gco='git checkout'          # Switch branches or restore files
alias gcon='git checkout -b'      # Checkout and create a new branch
alias gcm='git commit -m'         # Commit with a message inline
alias gundo='git reset --soft HEAD~1' # Undo the last commit (soft reset)
`

func InstallAliases(configFile string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Failed to get home directory:", err)
		return
	}
	fullPath := filepath.Join(homeDir, configFile)

	if utils.DryRun {
		fmt.Printf("[Dry Run] Would append the following aliases to %s:\n%s\n", fullPath, GitAliases)
		fmt.Println("[Dry Run] ✅ Aliases would be written to config file.")
		return
	}

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("❌ Failed to open config file:", err)
		return
	}
	defer f.Close()

	// Append the aliases
	_, err = f.WriteString("\n" + GitAliases + "\n")
	if err != nil {
		fmt.Println("❌ Failed to write to config file:", err)
		return
	}
}

func AskToInstallAliases(configFile string) bool {
	installAliases := true // Default to true for non-interactive mode or if user says yes
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to setup some git aliases?")
		if !answer {
			installAliases = false
		}
	} else {
		fmt.Println("Non-interactive mode: Setting up Git aliases by default.")
	}

	if !installAliases {
		fmt.Println("Skipping Git alias setup.")
		return false
	}

	InstallAliases(configFile)
	return true
}
