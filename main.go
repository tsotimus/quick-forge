package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tsotimus/quickforge/cmd"
	"github.com/tsotimus/quickforge/utils"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "quickforge",
		Short: "QuickForge sets up your Mac dev environment",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("🔧 Welcome to QuickForge!")
			shell, ok := utils.GetParentShell()
			if !ok {
				fmt.Println("Parent shell not found")
				return
			}

			shellConfigFile, ok := utils.GetShellConfigFile(shell)
			if !ok {
				fmt.Println("Shell config file not found")
				return
			}
			fmt.Println("Shell config file:", shellConfigFile)

			if !cmd.CheckBrew() {
				cmd.InstallBrew(shellConfigFile)
				utils.RestartShell(shellConfigFile)
				os.Exit(0)
			} else {
				cmd.AskToInstallGit()
				cmd.SetupSSHKey()
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
