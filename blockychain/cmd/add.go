package cmd

import (
	"errors"
	"time"

	internal "../internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <message>",
	Short: "Add a block to the blockchain.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires an argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		block := internal.NewBlock(time.Now().Format("02/01/2006"), internal.Data{Value: args[0]})
		blockchain.AddBlock(repository, block)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
