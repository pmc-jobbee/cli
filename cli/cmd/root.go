package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pmc-jobbee-cli",
	Short: "CLI tool for the Jobbee system",
	Long:  `CLI tool for the Jobbee system.`,
}

func init() {
	rootCmd.AddCommand(cmdVersion)
}

func Execute() error {
	return rootCmd.Execute()
}
