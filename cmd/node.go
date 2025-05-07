package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallFnm() {
	fmt.Println("🌐 Installing fnm...")

	// Run the install script silently
	cmd := exec.Command("sh", "-c", "curl -fsSL https://fnm.vercel.app/install | bash")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install fnm:", err)
		return
	}
}

func InstallNode() {
	fmt.Println("📦 Installing Node.js v22 via fnm...")

	cmd := exec.Command("fnm", "install", "22")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Node.js v22:", err)
		return
	}

	cmd = exec.Command("fnm", "use", "22")
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		fmt.Println("⚠️ Node installed, but couldn't activate version 22:", err)
		return
	}

	fmt.Println("✅ Node.js v22 installed and activated via fnm.")
}

func AskToInstallNode() {
	answer := ui.AskYesNo("📦 Do you want to install Node.js and fnm? (fast node manager)")
	if !answer {
		return
	}

	InstallFnm()
	InstallNode()
}
