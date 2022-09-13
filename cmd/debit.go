package cmd

import (
	"github.com/MrzBldk/accountant-cli/database"
	"github.com/spf13/cobra"
)

var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "The command that creates a credit transaction",
	Long:  "A command that creates a credit transaction for a particular user.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.AddTransaction(args[0], database.Debit, debitAmount, debitDescription)
	},
}

var debitDescription string
var debitAmount uint

func init() {
	rootCmd.AddCommand(debitCmd)
	debitCmd.Flags().StringVarP(&debitDescription, "description", "d", "", "Description for this debit transaction")
	debitCmd.Flags().UintVarP(&debitAmount, "amount", "a", 0, "Amount to be debited")
	debitCmd.MarkFlagRequired("description")
	debitCmd.MarkFlagRequired("amount")
	debitCmd.SetUsageTemplate("Usage: accountant-cli debit <username> --amount=<amount> --description==<description>")
}
