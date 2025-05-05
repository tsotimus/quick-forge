package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/tsotimus/quickforge/ui"
)

func InstallVSCode() {
	fmt.Println("Installing VSCode...")

	cmd := exec.Command("brew", "install", "--cask", "visual-studio-code")
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Suppress all output by redirecting to /dev/null
	nullDevice := "/dev/null"
	cmd.Stdout = exec.Command("tee", nullDevice).Stdout
	cmd.Stderr = exec.Command("tee", nullDevice).Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install VSCode.")
		return
	}

	fmt.Println("✅ Visual Studio Code installed successfully.")
}

func InstallCursor() {
	fmt.Println("Installing Cursor...")

	cmd := exec.Command("brew", "install", "--cask", "cursor")
	devNull, _ := os.Open(os.DevNull)
	defer devNull.Close()

	cmd.Stdout = devNull
	cmd.Stderr = devNull

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to install Cursor.")
		return
	}

	fmt.Println("✅ Cursor installed successfully.")
}

func InstallEditor() {
	answer := ui.AskSimpleChoice("Which editor do you want to install?", []string{"VSCode", "Cursor", "None"})
	if answer == "None" {
		fmt.Println("Skipping editor installation.")
		return
	}

	if answer == "VSCode" {
		InstallVSCode()
	} else if answer == "Cursor" {
		InstallCursor()
	}
}
