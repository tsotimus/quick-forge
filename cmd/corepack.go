package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
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
	answer := ui.AskYesNo("Do you want to install Corepack?")
	if !answer {
		return
	}

	InstallCorepack()
}
