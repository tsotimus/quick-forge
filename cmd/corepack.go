package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallCorepack() {
	fmt.Println("Installing Corepack...")

	cmd := exec.Command("brew", "install", "corepack")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Corepack:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Corepack installed successfully.")
}

func AskToInstallCorepack() {
	installCorepack := true
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to install Corepack?")
		if !answer {
			installCorepack = false
		}
	} else {
		fmt.Println("Non-interactive mode: Installing Corepack by default.")
	}

	if !installCorepack {
		fmt.Println("Skipping Corepack installation.")
		return
	}

	InstallCorepack()
}
