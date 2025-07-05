package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallWrap() {
	cmdToRun := []string{"brew", "install", "wrap"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install wrap with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ wrap would be installed successfully.") // Assuming success for dry run message
		return
	}
	fmt.Println("🔍 Installing wrap...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)
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
