package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallCorepack() {
	fmt.Println("🌐 Installing Corepack...")

	cmd := exec.Command("brew", "install", "corepack")

	// Silence both stdout and stderr
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Corepack:", err)
		return
	}

	fmt.Println("✅ Corepack installed successfully.")
}

func AskToInstallCorepack() {
	answer := ui.AskYesNo("Do you want to install Corepack?")
	if !answer {
		fmt.Println("🔕 Skipping Corepack installation.")
		return
	}

	InstallCorepack()
}
