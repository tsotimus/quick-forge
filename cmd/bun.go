package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallBum(shell string) {
	commandStr := fmt.Sprintf("curl -fsSL https://github.com/owenizedd/bum/raw/main/install.sh | %s", shell)
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Bum with command: sh -c \"%s\"\n", commandStr)
		fmt.Println("[Dry Run] ✅ Bum would be installed successfully.")
		return
	}

	fmt.Println("🌐 Installing Bum (Bun version manager)...")

	// Set up the command: sh -c "curl -fsSL <url> | bash"
	cmd := exec.Command("sh", "-c", commandStr)

	// Silence stdout and stderr
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Bum:", err)
		return
	}

	fmt.Println("✅ Bum installed successfully.")
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
	InstallBum(shell)

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
