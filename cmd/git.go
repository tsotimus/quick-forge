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
	answer := ui.AskYesNo("Do you want to install Git?")
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
	answer := ui.AskYesNo("Set up a GitHub SSH key?")
	if !answer {
		fmt.Println("🔕 Skipping SSH key setup.")
		return
	}

	fmt.Println("🔐 Setting up GitHub SSH key...")

	email := ui.AskForInput(
		"✉️  Enter your GitHub email address: ",
		"example@gmail.com",
		input.WithHelp(true),
		input.WithValidateFunc(validateEmail),
	)

	home := os.Getenv("HOME")
	sshKeyPath := home + "/.ssh/id_ed25519"
	sshConfigPath := home + "/.ssh/config"

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

	// Start SSH agent if not already running
	sshAuthSock := os.Getenv("SSH_AUTH_SOCK")
	sshAgentPid := os.Getenv("SSH_AGENT_PID")

	if sshAuthSock == "" || sshAgentPid == "" {
		fmt.Println("🚀 Starting ssh-agent...")
		out, err := exec.Command("ssh-agent", "-s").Output()
		if err != nil {
			fmt.Println("❌ Failed to start ssh-agent:", err)
			return
		}
		fmt.Print(string(out)) // Print agent info

		// Parse and set SSH_AUTH_SOCK and SSH_AGENT_PID
		reSock := regexp.MustCompile(`SSH_AUTH_SOCK=([^;]+);`)
		if match := reSock.FindStringSubmatch(string(out)); len(match) > 1 {
			sshAuthSock = match[1]
			os.Setenv("SSH_AUTH_SOCK", sshAuthSock)
		}
		rePid := regexp.MustCompile(`SSH_AGENT_PID=([0-9]+);`)
		if match := rePid.FindStringSubmatch(string(out)); len(match) > 1 {
			sshAgentPid = match[1]
			os.Setenv("SSH_AGENT_PID", sshAgentPid)
		}
	} else {
		fmt.Println("🧠 Using existing ssh-agent.")
	}

	// Ensure SSH config exists
	if _, err := os.Stat(sshConfigPath); os.IsNotExist(err) {
		_ = os.MkdirAll(home+"/.ssh", 0700)
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
