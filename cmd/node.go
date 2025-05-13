package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallFnm() {
	fmt.Println("🌐 Installing fnm...")

	// Run the install script and capture output
	cmd := exec.Command("sh", "-c", "curl -fsSL https://fnm.vercel.app/install | bash")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install fnm:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}
}

func InstallNode() {
	fmt.Println("📦 Installing Node.js v22 via fnm...")

	cmd := exec.Command("fnm", "install", "22")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Node.js v22:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	cmd = exec.Command("fnm", "use", "22")
	output, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Println("⚠️ Node installed, but couldn't activate version 22:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
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
