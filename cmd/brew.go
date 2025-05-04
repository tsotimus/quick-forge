package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CheckBrew() bool {
	_, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("🚨 Homebrew is not installed.")
		return false
	} else {
		fmt.Println("✅ Homebrew is installed.")
		return true
	}
}

func InstallBrew(shellConfigFile string) {
	fmt.Println("🛠️ Installing Homebrew...")

	cmd := exec.Command("/bin/bash", "-c", "curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash >/dev/null 2>&1")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Homebrew:", err)
		return
	}

	fmt.Println("✅ Homebrew installed. Updating shell config...")

	// Step 2: Add brew shellenv to ~/.zshrc or ~/.bashrc
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Could not determine home directory:", err)
		return
	}

	// Use .zshrc instead of .bashrc
	rcFile := filepath.Join(home, shellConfigFile)
	shellenvLine := `eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"`

	appendCmd := exec.Command("sh", "-c", fmt.Sprintf(`echo >> %s && echo '%s' >> %s`, rcFile, shellenvLine, rcFile))
	appendCmd.Stdout = os.Stdout
	appendCmd.Stderr = os.Stderr
	if err := appendCmd.Run(); err != nil {
		fmt.Println("❌ Failed to update shell config:", err)
		return
	}

	// Step 3: Export brew path for current session
	evalCmd := exec.Command("bash", "-c", shellenvLine)
	evalCmd.Stdout = os.Stdout
	evalCmd.Stderr = os.Stderr
	if err := evalCmd.Run(); err != nil {
		fmt.Println("⚠️ Brew installed but failed to update current PATH. You might need to restart your shell.")
		return
	}

	fmt.Println("✅ Homebrew path setup complete.")
}

func CheckBrewVersion() {
	cmd := exec.Command("brew", "--version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ Failed to check Homebrew version:", err)
		return
	}
	fmt.Printf("✅ Homebrew version:\n%s\n", output)
}
