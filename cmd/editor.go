package cmd

import (
	"fmt"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
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
	var editorToInstall string

	if !utils.NonInteractive {
		chosenEditor := ui.AskSimpleChoice("Which editor do you want to install?", []string{"VSCode", "Cursor", "None"})
		editorToInstall = chosenEditor
	} else {
		fmt.Println("Non-interactive mode: Defaulting to VSCode for editor installation.")
		editorToInstall = "VSCode"
	}

	if editorToInstall == "None" {
		fmt.Println("Skipping editor installation.")
		return
	}

	if editorToInstall == "VSCode" {
		InstallVSCode()
	} else if editorToInstall == "Cursor" {
		InstallCursor()
	}
}
