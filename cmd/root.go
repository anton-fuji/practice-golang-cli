package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-cli",
	Short: "A CLI application with Bubble Tea",
	Long:  `This is a CLI tool built with Cobra and Bubble Tea for interactive TUI applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'golang-cli tui' to launch the TUI.")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
