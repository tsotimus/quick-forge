package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallBun(shell string) {
	commandStr := fmt.Sprintf("curl -fsSL https://github.com/owenizedd/bum/raw/main/install.sh | %s", shell)
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Bun via Bum with command: sh -c \"%s\"\n", commandStr)
		fmt.Println("[Dry Run] ✅ Bun would be installed successfully via Bum.")
		return
	}

	fmt.Println("🌐 Installing Bun via Bum...")

	// Set up the command: sh -c "curl -fsSL <url> | bash"
	cmd := exec.Command("sh", "-c", commandStr)

	// Silence stdout and stderr
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Bun:", err)
		return
	}

	fmt.Println("✅ Bun installed successfully via Bum.")
}

func AskToInstallBun(shell string) {
	installBun := true // Default to true for non-interactive mode or if user says yes
	if !utils.NonInteractive {
		answer := ui.AskYesNo("Do you want to install Bun?")
		if !answer {
			installBun = false
		}
	} else {
		fmt.Println("Non-interactive mode: Installing Bun by default.")
	}

	if !installBun {
		fmt.Println("Skipping Bun installation.")
		return
	}
	InstallBun(shell)
}
