package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallBun() {
	fmt.Println("🌐 Installing Bun via Bum...")

	// Set up the command: sh -c "curl -fsSL <url> | bash"
	cmd := exec.Command("sh", "-c", "curl -fsSL https://github.com/owenizedd/bum/raw/main/install.sh | bash")

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

func AskToInstallBun() {
	answer := ui.AskYesNo("Do you want to install Bun?")
	if !answer {
		fmt.Println("🔕 Skipping Bun installation.")
		return
	}
	InstallBun()
}
