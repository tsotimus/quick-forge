package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallWrap() {
	fmt.Println("🔍 Installing wrap...")

	cmd := exec.Command("brew", "install", "wrap")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install wrap:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}
}

func AskToInstallWrap() {
	answer := ui.AskYesNo("Do you want to install wrap?")
	if answer {
		InstallWrap()
	}
}
