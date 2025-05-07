package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/tsotimus/quickforge/cmd"
	"github.com/tsotimus/quickforge/utils"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "quickforge",
		Short: "QuickForge sets up your Mac dev environment",
		Run: func(_ *cobra.Command, args []string) {
			lightning := lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render("⚡")
			hammer := lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Render("🔨")
			fmt.Println(lightning + " Welcome to QuickForge! " + hammer)
			utils.CheckOSSupported()
			shell, ok := utils.DetectShell()
			if !ok {
				fmt.Println("Shell not found")
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
				cmd.AskToInstallEditor()
				cmd.AskToInstallNode()
				cmd.AskToInstallCorepack()
				cmd.AskToInstallBun(shell)
				cmd.AskToInstallAliases(shellConfigFile)
				cmd.AskToInstallBrowsers()
				utils.Finish(shellConfigFile)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
