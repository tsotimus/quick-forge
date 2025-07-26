package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-ps"
)

func DetectShell() (string, bool) {
	// First try $SHELL
	if shellPath := os.Getenv("SHELL"); shellPath != "" {
		return filepath.Base(shellPath), true
	}

	// Fallback to parent process
	pid := os.Getppid()
	process, err := ps.FindProcess(pid)
	if err == nil && process != nil {
		return process.Executable(), true
	}

	// Couldn't detect
	return "", false
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

func RestartShellForNode(shellConfigFile string) {
	fmt.Println("\nfnm has been installed! Please restart your shell to continue with Node.js installation:")
	fmt.Printf("  source ~/%s && quickforge --resume-from-node\n", shellConfigFile)
}

func RestartShellForBun(shellConfigFile string) {
	fmt.Println("\nbum has been installed! Please restart your shell to continue with Bun installation:")
	fmt.Printf("  source ~/%s && quickforge --resume-from-bun\n", shellConfigFile)
}

func Finish(shellConfigFile string, restartShell bool) {
	fmt.Println("\nAll done!")
	if restartShell {
		fmt.Println("\nRestart your shell to apply all changes.")
		fmt.Printf("  source ~/%s\n", shellConfigFile)
	}
}
