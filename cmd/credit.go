package cmd

import (
	"github.com/MrzBldk/accountant-cli/database"
	"github.com/spf13/cobra"
)

var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "A command that creates a credit transaction",
	Long:  "A command that creates a credit transaction for a particular user.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.AddTransaction(args[0], database.Credit, creditAmount, creditDescription)
	},
}

var creditDescription string
var creditAmount uint

func init() {
	rootCmd.AddCommand(creditCmd)
	creditCmd.Flags().StringVarP(&creditDescription, "description", "d", "", "Description for this credit transaction")
	creditCmd.Flags().UintVarP(&creditAmount, "amount", "a", 0, "Amount to be credited")
	creditCmd.MarkFlagRequired("description")
	creditCmd.MarkFlagRequired("amount")
	creditCmd.SetUsageTemplate("Usage: accountant-cli credit <username> --amount=<amount> --description==<description>")
}
