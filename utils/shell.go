package utils

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

func GetParentShell() (string, bool) {
	pid := os.Getppid()
	process, err := ps.FindProcess(pid)

	if err != nil || process == nil {
		fmt.Println("Could not detect parent process.")
		return "", false
	}
	fmt.Println("Parent process:", process.Executable())
	return process.Executable(), true
}

// Input: bash, zsh, fish, etc
func GetShellConfigFile(shell string) (string, bool) {
	switch shell {
	case "bash":
		return ".bashrc", true
	case "zsh":
		return ".zshrc", true
	case "fish":
		return ".config/fish/config.fish", true
	default:
		return "", false
	}
}

func RestartShell(shellConfigFile string) {
	fmt.Println("\nNext step:")
	fmt.Printf("  source ~/%s && quickforge\n", shellConfigFile)
}
