package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallChrome() {
	cmdToRun := []string{"brew", "install", "--cask", "google-chrome"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Google Chrome with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Google Chrome would be installed successfully.")
		return
	}
	fmt.Println("🔍 Installing Google Chrome...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

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
	cmdToRun := []string{"brew", "install", "--cask", "zen-browser"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Zen Browser with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Zen Browser would be installed successfully.")
		return
	}
	fmt.Println("🔍 Installing Zen Browser...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

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
	cmdToRun := []string{"brew", "install", "--cask", "arc-browser"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Arc Browser with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Arc Browser would be installed successfully.")
		return
	}
	fmt.Println("🔍 Installing Arc Browser...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

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
	var browsersToInstall []string

	if !utils.NonInteractive {
		browsersToInstall = ui.AskMultiChoice("Which browsers do you want to install?", []string{"Google Chrome", "Zen Browser", "Arc Browser", "None"})
	} else {
		fmt.Println("Non-interactive mode: Defaulting to install Google Chrome.")
		browsersToInstall = []string{"Google Chrome"}
	}

	// Filter out "None" option if selected
	var filteredBrowsers []string
	for _, browser := range browsersToInstall {
		if browser != "None" {
			filteredBrowsers = append(filteredBrowsers, browser)
		}
	}
	browsersToInstall = filteredBrowsers

	if len(browsersToInstall) == 0 {
		fmt.Println("Skipping browser installation.")
		return
	}

	for _, browser := range browsersToInstall {
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
