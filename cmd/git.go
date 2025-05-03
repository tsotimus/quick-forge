package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func AskToInstallGit() {
	fmt.Print("Do you want to install Git? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" || input == "yes" {
		fmt.Println("Installing Git with Homebrew...")
		cmd := exec.Command("brew", "install", "git")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("❌ Failed to install Git:", err)
		} else {
			fmt.Println("✅ Git installed successfully.")
		}
	} else {
		fmt.Println("❌ Git will not be installed.")
	}
}
