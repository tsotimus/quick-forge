package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallBun() {
	fmt.Println("🌐 Installing Bun...")

	cmd := exec.Command("brew", "install", "bun")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Bun:", err)
		return
	}

}
func AskToInstallBun() {
	answer := ui.AskYesNo("Do you want to install Bun?")
	if !answer {
		fmt.Println("🔕 Skipping Bun installation.")
		return
	}

}
