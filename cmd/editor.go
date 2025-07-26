package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/tsotimus/quickforge/ui"
	"github.com/tsotimus/quickforge/utils"
)

func InstallVSCode() {
	cmdToRun := []string{"brew", "install", "--cask", "visual-studio-code"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install VSCode with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Visual Studio Code would be installed successfully.")
		return
	}
	fmt.Println("Installing VSCode...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

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
	cmdToRun := []string{"brew", "install", "--cask", "cursor"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Cursor with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Cursor would be installed successfully.")
		return
	}
	fmt.Println("Installing Cursor...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

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

func InstallZed() {
	cmdToRun := []string{"brew", "install", "--cask", "zed"}
	if utils.DryRun {
		fmt.Printf("[Dry Run] Would install Zed with command: %s\n", strings.Join(cmdToRun, " "))
		fmt.Println("[Dry Run] ✅ Zed would be installed successfully.")
		return
	}
	fmt.Println("Installing Zed...")
	cmd := exec.Command(cmdToRun[0], cmdToRun[1:]...)

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ Failed to install Zed:", err)
		fmt.Println("--- Command output ---")
		fmt.Println(string(output))
		fmt.Println("----------------------")
		return
	}

	fmt.Println("✅ Zed installed successfully.")
}

func AskToInstallEditor() {
	var editorToInstall string

	if !utils.NonInteractive {
		chosenEditor := ui.AskSimpleChoice("Which editor do you want to install?", []string{"VSCode", "Cursor", "Zed", "None"})
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
	} else if editorToInstall == "Zed" {
		InstallZed()
	}
}
