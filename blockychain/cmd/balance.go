package cmd

import (
	"errors"
	"fmt"

	internal "../internal"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance <address>",
	Short: "Get balance of an address.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		blockchain := internal.GetBlockchain(repository, difficulty)
		fmt.Printf("Balance of '%s': %d\n", args[0], blockchain.GetBalance(repository, args[0]))
	},
}

func init() {
	RootCmd.AddCommand(balanceCmd)
}
