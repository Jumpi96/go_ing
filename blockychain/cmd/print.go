package cmd

import (
	"fmt"
	"strconv"

	internal "../internal"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print <message>",
	Short: "Print the blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		blockchain := internal.GetBlockchain(repository, difficulty)
		iterator := blockchain.Iterator(repository)
		block := iterator.Next()
		for block != nil {
			fmt.Printf("Prev. hash: %x\n", block.PreviousHash)
			fmt.Printf("Data: %v\n", block.Transactions)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()

			block = iterator.Next()
		}
		fmt.Printf("PoW: %s\n", strconv.FormatBool(blockchain.IsChainValid(repository)))
	},
}

func init() {
	RootCmd.AddCommand(printCmd)
}
