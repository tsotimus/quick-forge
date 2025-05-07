package cmd

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallChrome() {
	fmt.Println("🔍 Installing Google Chrome...")
	cmd := exec.Command("brew", "install", "--cask", "google-chrome")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Google Chrome:", err)
		return
	}

	fmt.Println("✅ Google Chrome installed successfully.")
}

func InstallZen() {
	fmt.Println("🔍 Installing Zen Browser...")
	cmd := exec.Command("brew", "install", "--cask", "zen-browser")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Zen Browser:", err)
		return
	}

	fmt.Println("✅ Zen Browser installed successfully.")
}

func InstallArc() {
	fmt.Println("🔍 Installing Arc Browser...")
	cmd := exec.Command("brew", "install", "--cask", "arc-browser")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Arc Browser:", err)
		return
	}

	fmt.Println("✅ Arc Browser installed successfully.")
}

func AskToInstallBrowsers() {
	answer := ui.AskMultiChoice("Which browsers do you want to install?", []string{"Google Chrome", "Zen Browser", "Arc Browser"})
	if len(answer) == 0 {
		fmt.Println("🔕 Skipping browser installation.")
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
