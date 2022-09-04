package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "accountant-cli",
	Short: "A CLI that helps you manage user accounts",
	Long: `A command-line interface that helps you manage user accounts.
You can add credit and debit transactions to user accounts.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
