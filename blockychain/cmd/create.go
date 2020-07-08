package cmd

import (
	"fmt"
	"log"

	internal "../internal"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		blockchain := internal.GetBlockchain(repository, difficulty)
		if blockchain != nil {
			log.Panic("A blockchain already exists!")
		} else {
			internal.NewBlockchain(repository, address, difficulty)
			fmt.Println("Done!")
		}
	},
}

var address string

func init() {
	createCmd.Flags().StringVarP(&address, "address", "a", "jumposhi", "address for the initial coin")

	RootCmd.AddCommand(createCmd)
}
