package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallChrome() {
	fmt.Println("🔍 Installing Google Chrome...")
	cmd := exec.Command("brew", "install", "--cask", "google-chrome")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Google Chrome:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Google Chrome installed successfully.")
}

func InstallZen() {
	fmt.Println("🔍 Installing Zen Browser...")
	cmd := exec.Command("brew", "install", "--cask", "zen-browser")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Zen Browser:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Zen Browser installed successfully.")
}

func InstallArc() {
	fmt.Println("🔍 Installing Arc Browser...")
	cmd := exec.Command("brew", "install", "--cask", "arc-browser")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Arc Browser:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Arc Browser installed successfully.")
}

func AskToInstallBrowsers() {
	answer := ui.AskMultiChoice("Which browsers do you want to install?", []string{"Google Chrome", "Zen Browser", "Arc Browser"})
	if len(answer) == 0 {
		return
	}

	for _, browser := range answer {
		switch browser {
		case "Google Chrome":
			InstallChrome()
		case "Zen Browser":
			InstallZen()
		case "Arc Browser":
			InstallArc()
		}
	}
}
