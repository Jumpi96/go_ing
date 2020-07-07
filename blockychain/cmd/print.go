package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print <message>",
	Short: "Print the blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		iterator := blockchain.Iterator(repository)
		block := iterator.Next()
		for block != nil {
			fmt.Printf("Prev. hash: %x\n", block.PreviousHash)
			fmt.Printf("Data: %s\n", block.Data)
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
