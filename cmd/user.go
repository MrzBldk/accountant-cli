package cmd

import (
	"fmt"

	"github.com/MrzBldk/accountant-cli/database"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A command that displays user information",
	Long:  "A command that displays user name, balance and all his transactions",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetUser(args[0])
		fmt.Println(user.Name, user.Balance)
		for _, v := range user.Transactions {
			fmt.Println(v.Type, v.Amount, v.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.SetUsageTemplate("Usage: accountant-cli user <username>")
}
