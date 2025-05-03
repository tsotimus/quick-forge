package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tsotimus/quickforge/cmd"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "quickforge",
		Short: "QuickForge sets up your Mac dev environment",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("🔧 Welcome to QuickForge!")
			if cmd.CheckBrew() {
				cmd.AskToInstallGit()
			} else {
				cmd.InstallBrew()
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
