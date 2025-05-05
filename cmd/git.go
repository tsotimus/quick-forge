package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/cqroot/prompt/input"
	"github.com/tsotimus/quickforge/ui"
)

var ErrInvalidEmail = errors.New("invalid email address")

func validateEmail(email string) error {
	// Simple regex for demonstration; you can use a more robust one if needed
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return fmt.Errorf("%s: %w", email, ErrInvalidEmail)
	}
	return nil
}

func AskToInstallGit() {
	answer := ui.AskYesNo("Do you want to install Git? (y/n): ")
	if !answer {
		fmt.Println("❌ Git will not be installed.")
		return
	}

	fmt.Println("Installing Git with Homebrew...")
	var outBuf, errBuf strings.Builder
	cmd := exec.Command("brew", "install", "git")
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Failed to install Git:", err)
		if outBuf.Len() > 0 {
			fmt.Println("--- brew output ---")
			fmt.Print(outBuf.String())
		}
		if errBuf.Len() > 0 {
			fmt.Println("--- brew error output ---")
			fmt.Print(errBuf.String())
		}
	} else {
		fmt.Println("✅ Git installed successfully.")
	}
}

func SetupSSHKey() {
	answer := ui.AskYesNo("📦 Set up a GitHub SSH key? (y/n): ")
	if !answer {
		fmt.Println("🔕 Skipping SSH key setup.")
		return
	}

	fmt.Println("🔐 Setting up GitHub SSH key...")

	// Get GitHub email
	email := ui.AskForInput(
		"✉️  Enter your GitHub email address: ",
		"example@gmail.com",
		input.WithHelp(true),
		input.WithValidateFunc(validateEmail),
	)

	sshKeyPath := os.Getenv("HOME") + "/.ssh/id_ed25519"
	sshConfigPath := os.Getenv("HOME") + "/.ssh/config"

	// Generate SSH key
	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", sshKeyPath, "-N", "")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to generate SSH key:", err)
		return
	}
	fmt.Println("✅ SSH key generated.")

	// Start SSH agent
	out, err := exec.Command("ssh-agent", "-s").Output()
	if err != nil {
		fmt.Println("❌ Failed to start ssh-agent:", err)
		return
	}
	fmt.Print(string(out)) // Show agent pid

	// Ensure SSH config exists
	if _, err := os.Stat(sshConfigPath); os.IsNotExist(err) {
		_ = os.MkdirAll(os.Getenv("HOME")+"/.ssh", 0700)
		_, _ = os.Create(sshConfigPath)
	}

	// Check if config already contains github.com
	configBytes, _ := os.ReadFile(sshConfigPath)
	if !strings.Contains(string(configBytes), "Host github.com") {
		f, _ := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_WRONLY, 0600)
		defer f.Close()
		sshConfig := `
Host github.com
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_ed25519
`
		_, _ = f.WriteString(sshConfig)
		fmt.Println("🛠️  SSH config updated.")
	}

	// Add key to agent
	addCmd := exec.Command("ssh-add", sshKeyPath)
	var addOut, addErr strings.Builder
	addCmd.Stdout = &addOut
	addCmd.Stderr = &addErr
	if err := addCmd.Run(); err != nil {
		fmt.Println("❌ Failed to add key to ssh-agent:", err)
		if addOut.Len() > 0 {
			fmt.Println("--- ssh-add output ---")
			fmt.Print(addOut.String())
		}
		if addErr.Len() > 0 {
			fmt.Println("--- ssh-add error output ---")
			fmt.Print(addErr.String())
		}
	} else {
		fmt.Println("🔑 SSH key added to ssh-agent.")
	}

	// Output public key
	pubKey, err := os.ReadFile(sshKeyPath + ".pub")
	if err != nil {
		fmt.Println("❌ Failed to read public key:", err)
		return
	}
	fmt.Println("\n📋 Public key:")
	fmt.Println(string(pubKey))
	fmt.Println("\n🌐 Add this key to your GitHub account:")
	fmt.Println("👉 https://github.com/settings/keys")
}
