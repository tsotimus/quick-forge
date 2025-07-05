package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cqroot/prompt/input"
	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
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
	installGit := true
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to install Git?")
		if !answer {
			installGit = false
		}
	}

	if !installGit {
		fmt.Println("Skipping Git installation.")
		return
	}

	gitInstallCmdParts := []string{"brew", "install", "git"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Git with command: %s\n", strings.Join(gitInstallCmdParts, " "))
		fmt.Println("[Dry Run] ✅ Git would be installed successfully.")
		// We don't return here if we want SetupSSHKey to also show its dry-run steps
		// However, AskToInstallGit is usually followed by SetupSSHKey in main.go, so this is fine.
		// For now, if Git isn't "installed" (even in dry run), SetupSSHKey might not make sense to show.
		// Let's assume for dry run, if we "would install Git", we also proceed to "would set up SSH key".
	} else {
		fmt.Println("Installing Git with Homebrew...")
		var outBuf, errBuf strings.Builder
		cmd := exec.Command(gitInstallCmdParts[0], gitInstallCmdParts[1:]...)
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
			return // If actual install fails, stop here
		} else {
			fmt.Println("✅ Git installed successfully.")
		}
	}
}

func SetupSSHKey() {
	setupKey := true
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Set up a GitHub SSH key?")
		if !answer {
			setupKey = false
		}
	}

	if !setupKey {
		// No dry run message needed if user/default skips entirely
		return
	}

	var email string
	if !utils.NonInteractive {
		email = ui.AskForInput(
			"✉️  Enter your GitHub email address: ",
			"example@gmail.com",
			input.WithHelp(true),
			input.WithValidateFunc(validateEmail),
		)
	} else {
		cmdEmail := exec.Command("git", "config", "--global", "user.email")
		emailBytes, err := cmdEmail.Output() // This read operation is not part of dry run itself
		if err == nil && len(strings.TrimSpace(string(emailBytes))) > 0 {
			email = strings.TrimSpace(string(emailBytes))
			fmt.Printf("✉️ Using Git config email: %s\n", email)
		} else {
			fmt.Println("⚠️ Git user email not found in global config. Cannot proceed with SSH key generation without an email.")
			fmt.Println("Skipping GitHub SSH key setup.")
			return
		}
	}

	home := os.Getenv("HOME")
	sshKeyPath := home + "/.ssh/id_ed25519"
	sshKeyPubPath := sshKeyPath + ".pub"
	sshConfigPath := home + "/.ssh/config"
	sshKeygenCmdParts := []string{"ssh-keygen", "-t", "ed25519", "-C", email, "-f", sshKeyPath, "-N", ""}
	sshAgentStartCmd := "ssh-agent -s" // Actual command might be more complex to parse output from
	sshAddCmdParts := []string{"ssh-add", sshKeyPath}
	sshConfigContent := `
Host github.com
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_ed25519
`

	if utils.DryRun {
		fmt.Printf("[Dry Run] 🔐 Would set up GitHub SSH key for email: %s\n", email)
		fmt.Printf("[Dry Run] Would generate SSH key with command: %s\n", strings.Join(sshKeygenCmdParts, " "))
		fmt.Printf("[Dry Run] Would ensure SSH agent is running (e.g., by trying to execute: %s and parsing output)\n", sshAgentStartCmd)
		fmt.Printf("[Dry Run] Would ensure SSH config directory exists: %s/.ssh\n", home)
		fmt.Printf("[Dry Run] Would ensure SSH config file exists: %s\n", sshConfigPath)
		fmt.Printf("[Dry Run] Would check if '%s' contains 'Host github.com' and append if not.\n", sshConfigPath)
		fmt.Printf("[Dry Run] Content to append to SSH config (if needed):\n%s\n", sshConfigContent)
		fmt.Printf("[Dry Run] Would add SSH key to agent with command: %s\n", strings.Join(sshAddCmdParts, " "))
		fmt.Printf("[Dry Run] Would read public key from: %s\n", sshKeyPubPath)
		fmt.Println("[Dry Run] ✅ SSH key setup would be complete.")
		return
	}

	fmt.Println("🔐 Setting up GitHub SSH key...")

	// Generate SSH key
	cmdSshKeygen := exec.Command(sshKeygenCmdParts[0], sshKeygenCmdParts[1:]...)
	cmdSshKeygen.Stdin = os.Stdin
	cmdSshKeygen.Stdout = os.Stdout
	cmdSshKeygen.Stderr = os.Stderr
	if err := cmdSshKeygen.Run(); err != nil {
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
		fmt.Print(string(out))
		reSock := regexp.MustCompile(`SSH_AUTH_SOCK=([^;]+);`)
		if match := reSock.FindStringSubmatch(string(out)); len(match) > 1 {
			os.Setenv("SSH_AUTH_SOCK", match[1])
		}
		rePid := regexp.MustCompile(`SSH_AGENT_PID=([0-9]+);`)
		if match := rePid.FindStringSubmatch(string(out)); len(match) > 1 {
			os.Setenv("SSH_AGENT_PID", match[1])
		}
	} else {
		fmt.Println("🧠 Using existing ssh-agent.")
	}

	// Ensure SSH config exists and add github.com entry
	if _, err := os.Stat(sshConfigPath); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(sshConfigPath), 0700)
		_, _ = os.Create(sshConfigPath)
	}
	configBytes, _ := os.ReadFile(sshConfigPath)
	if !strings.Contains(string(configBytes), "Host github.com") {
		f, _ := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_WRONLY, 0600)
		defer f.Close()
		_, _ = f.WriteString(sshConfigContent)
		fmt.Println("🛠️  SSH config updated.")
	}

	// Add key to agent
	addCmd := exec.Command(sshAddCmdParts[0], sshAddCmdParts[1:]...)
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

	// Output public key (read operation, safe for dry run, but follows the pattern)
	pubKey, err := os.ReadFile(sshKeyPubPath)
	if err != nil {
		fmt.Println("❌ Failed to read public key:", err)
		return
	}
	fmt.Println("\n📋 Public key:")
	fmt.Println(string(pubKey))
	fmt.Println("\n🌐 Add this key to your GitHub account:")
	fmt.Println("👉 https://github.com/settings/keys")
}
