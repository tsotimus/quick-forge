package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/tsotimus/quickforge/cmd"
	"github.com/tsotimus/quickforge/utils"
)

var version = "dev" // This will be set by build flags in CI

func main() {
	var rootCmd = &cobra.Command{
		Use:     "quickforge",
		Short:   "QuickForge sets up your Typescript dev environment",
		Version: version,
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

			// Handle resume from specific steps
			if utils.ResumeFromNode {
				cmd.AskToInstallNode()
				cmd.AskToInstallCorepack()
				cmd.AskToInstallBun(shell)
				shouldRestart := cmd.AskToInstallAliases(shellConfigFile)
				cmd.AskToInstallBrowsers()
				cmd.AskToInstallWrap()
				utils.Finish(shellConfigFile, shouldRestart)
				return
			}

			if utils.ResumeFromBun {
				cmd.AskToInstallBun(shell)
				shouldRestart := cmd.AskToInstallAliases(shellConfigFile)
				cmd.AskToInstallBrowsers()
				cmd.AskToInstallWrap()
				utils.Finish(shellConfigFile, shouldRestart)
				return
			}

			// Normal flow
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
				shouldRestart := cmd.AskToInstallAliases(shellConfigFile)
				cmd.AskToInstallBrowsers()
				cmd.AskToInstallWrap()
				utils.Finish(shellConfigFile, shouldRestart)
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&utils.NonInteractive, "non-interactive", "y", false, "Enable non-interactive mode (accepts all defaults)")
	rootCmd.PersistentFlags().BoolVarP(&utils.DryRun, "dry-run", "d", false, "Simulate changes without executing them")
	rootCmd.PersistentFlags().BoolVar(&utils.ResumeFromNode, "resume-from-node", false, "Resume installation from Node.js step (after fnm is installed)")
	rootCmd.PersistentFlags().BoolVar(&utils.ResumeFromBun, "resume-from-bun", false, "Resume installation from Bun step (after bum is installed)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
