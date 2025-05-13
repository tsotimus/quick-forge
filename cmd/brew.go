package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tsotimus/quickforge/utils"
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
	// Define command strings for dry run output
	brewInstallCmdString := "/bin/bash -c \"curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash >/dev/null 2>&1\"" // Note: actual command redirects output

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Could not determine home directory:", err)
		return
	}
	rcFile := filepath.Join(home, shellConfigFile)
	// IMPORTANT: The path to brew shellenv might differ (e.g., /opt/homebrew/bin/brew for Apple Silicon)
	// This is a placeholder. A more robust solution would find brew's path after installation.
	shellenvLine := `eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"`
	appendShellCmdString := fmt.Sprintf(`echo >> %s && echo '%s' >> %s`, rcFile, shellenvLine, rcFile)
	evalShellCmdString := shellenvLine

	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Homebrew with command: %s\n", brewInstallCmdString)
		fmt.Println("[Dry Run] ✅ Homebrew installation and setup would be complete.")
		return
	}

	fmt.Println("🛠️ Installing Homebrew...")
	// Actual Homebrew installation
	cmd := exec.Command("/bin/bash", "-c", "curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash >/dev/null 2>&1")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Homebrew:", err)
		return
	}

	fmt.Println("✅ Homebrew installed. Updating shell config...")

	// Add brew shellenv to rcFile
	appendCmd := exec.Command("sh", "-c", appendShellCmdString)
	appendCmd.Stdout = os.Stdout
	appendCmd.Stderr = os.Stderr
	if err := appendCmd.Run(); err != nil {
		fmt.Println("❌ Failed to update shell config:", err)
		return
	}

	// Export brew path for current session
	evalCmd := exec.Command("bash", "-c", evalShellCmdString) // Use evalShellCmdString here
	evalCmd.Stdout = os.Stdout
	evalCmd.Stderr = os.Stderr
	if err := evalCmd.Run(); err != nil {
		fmt.Println("⚠️ Brew installed but failed to update current PATH. You might need to restart your shell.")
		return
	}

	fmt.Println("✅ Homebrew path setup complete.")
}
