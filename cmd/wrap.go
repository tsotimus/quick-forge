package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
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
	installWrap := true
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to install wrap?")
		if !answer {
			installWrap = false
		}
	} else {
		fmt.Println("Non-interactive mode: Installing Warp by default.")
	}

	if !installWrap {
		fmt.Println("Skipping Warp installation.")
		return
	}
	InstallWrap()
}
