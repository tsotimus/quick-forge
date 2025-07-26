package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallBum(shell string) bool {
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Bum with command: curl -fsSL https://github.com/owenizedd/bum/raw/main/install.sh | bash\n")
		fmt.Println("[Dry Run] ✅ Bum would be installed successfully.")
		return true
	}

	fmt.Println("🌐 Installing Bum (Bun version manager)...")

	// Use bash explicitly instead of the detected shell to avoid substitution issues
	cmd := exec.Command("bash", "-c", "curl -fsSL https://github.com/owenizedd/bum/raw/main/install.sh | bash")

	// Capture stdout and stderr to show detailed error information
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Bum:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return false
	}

	fmt.Println("✅ Bum installed successfully.")
	return true
}

func InstallBun(shell string) {
	if utils.DryRun {
		fmt.Println("[Dry Run] Would install Bun via Bum")
		fmt.Println("[Dry Run] ✅ Bun would be installed successfully via Bum.")
		return
	}

	fmt.Println("🌐 Installing Bun via Bum...")

	// Use bum to install the latest Bun
	cmd := exec.Command("bum", "install", "latest")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Bun via Bum:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Bun installed successfully via Bum.")
}

func AskToInstallBun(shell string) {
	installBun := true // Default to true for non-interactive mode or if user says yes
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to install Bun and Bum? (Bun version manager)")
		if !answer {
			installBun = false
		}
	} else {
		fmt.Println("Non-interactive mode: Installing Bun and Bum by default.")
	}

	if !installBun {
		fmt.Println("Skipping Bun installation.")
		return
	}

	// If we're resuming from Bun step, bum is already installed
	if utils.ResumeFromBun {
		InstallBun(shell)
		return
	}

	// Normal flow: install bum first, then exit for shell restart
	success := InstallBum(shell)
	if !success {
		fmt.Println("❌ Bum installation failed. Cannot proceed with Bun installation.")
		return
	}

	// Get shell config for restart instruction
	shellDetected, ok := utils.DetectShell()
	if !ok {
		fmt.Println("Shell not found")
		return
	}
	shellConfigFile, ok := utils.GetShellConfigFile(shellDetected)
	if !ok {
		fmt.Println("Shell config file not found")
		return
	}

	utils.RestartShellForBun(shellConfigFile)
	os.Exit(0)
}
