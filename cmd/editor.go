package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallVSCode() {
	fmt.Println("Installing VSCode...")

	cmd := exec.Command("brew", "install", "--cask", "visual-studio-code")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install VSCode:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Visual Studio Code installed successfully.")
}

func InstallCursor() {
	fmt.Println("Installing Cursor...")

	cmd := exec.Command("brew", "install", "--cask", "cursor")

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Cursor:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Cursor installed successfully.")
}

func AskToInstallEditor() {
	answer := ui.AskSimpleChoice("Which editor do you want to install?", []string{"VSCode", "Cursor", "None"})
	if answer == "None" {
		return
	}

	if answer == "VSCode" {
		InstallVSCode()
	} else if answer == "Cursor" {
		InstallCursor()
	}
}
