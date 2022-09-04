package cmd

import (
	"github.com/MrzBldk/accountant-cli/database"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A command that creates a user",
	Long:  `A command that creates a user with the specified name and balance.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.CreateUser(args[0], userBalance)
	},
}

var userBalance uint

func init() {
	userCmd.AddCommand(createCmd)
	createCmd.Flags().UintVarP(&userBalance, "balance", "b", 0, "User balance")
	createCmd.MarkFlagRequired("balance")
	createCmd.SetUsageTemplate("Usage: user create <username> --balance=<balance>")
}
