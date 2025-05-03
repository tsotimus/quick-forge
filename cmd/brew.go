package cmd

import (
	"fmt"
	"os/exec"
)

func CheckBrew() {
	_, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("🚨 Homebrew is not installed.")
	} else {
		fmt.Println("✅ Homebrew is installed.")
	}
}
