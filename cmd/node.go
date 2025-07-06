package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallFnm() {
	fnmInstallCmdStr := "sh -c \"curl -fsSL https://fnm.vercel.app/install | bash\""
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install fnm with command: %s\n", fnmInstallCmdStr)
		fmt.Println("[Dry Run] ✅ fnm installation script would be executed.")
		return
	}
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

	fmt.Println("✅ fnm installed successfully")
}

func InstallNode() {
	if utils.DryRun {
		fmt.Println("[Dry Run] Would install Node.js v22 via fnm with eval environment setup")
		fmt.Println("[Dry Run] Would run: bash -c \"eval $(~/.local/share/fnm/fnm env) && fnm install 22\"")
		fmt.Println("[Dry Run] Would run: bash -c \"eval $(~/.local/share/fnm/fnm env) && fnm use 22\"")
		fmt.Println("[Dry Run] ✅ Node.js v22 would be installed and activated via fnm.")
		return
	}

	fmt.Println("📦 Installing Node.js v22 via fnm...")

	// Use fnm's eval command to set up environment and install Node
	installCmd := "bash -c \"eval $(~/.local/share/fnm/fnm env) && fnm install 22\""
	cmd := exec.Command("sh", "-c", installCmd)
	outputInstall, errInstall := cmd.CombinedOutput()

	if errInstall != nil {
		fmt.Println("❌ Failed to install Node.js v22:", errInstall)
		fmt.Println("--- Command output ---")
		fmt.Println(string(outputInstall))
		fmt.Println("----------------------")
		return
	}

	// Use fnm's eval command to set up environment and activate Node
	useCmd := "bash -c \"eval $(~/.local/share/fnm/fnm env) && fnm use 22\""
	cmdUse := exec.Command("sh", "-c", useCmd)
	outputUse, errUse := cmdUse.CombinedOutput()

	if errUse != nil {
		fmt.Println("⚠️ Node installed, but couldn't activate version 22:", errUse)
		fmt.Println("--- Command output ---")
		fmt.Println(string(outputUse))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Node.js v22 installed and activated via fnm.")
}

func AskToInstallNode() {
	installNode := true // Default to true for non-interactive mode or if user says yes
	if !utils.NonInteractive {
		answer := ui.AskYesNo("📦 Do you want to install Node.js and fnm? (fast node manager)")
		if !answer {
			installNode = false
		}
	} else {
		fmt.Println("Non-interactive mode: Installing Node.js and fnm by default.")
	}

	if !installNode {
		fmt.Println("Skipping Node.js and fnm installation.")
		return
	}

	InstallFnm()
	InstallNode()
}
