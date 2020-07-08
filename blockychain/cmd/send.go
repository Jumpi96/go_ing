package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	internal "../internal"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send <amount>",
	Short: "Send coins through the blockchain.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires an argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		blockchain := internal.GetBlockchain(repository, difficulty)
		if amount, err := strconv.Atoi(args[0]); err == nil {
			tx := internal.NewUTXOTransaction(repository, blockchain, from, to, amount)
			block := internal.NewBlock(time.Now().Format("02/01/2006"), []*internal.Transaction{tx})
			blockchain.AddBlock(repository, block)
			fmt.Println("Success!")
		} else {
			log.Panic("Amount is not a correct integer.")
		}
	},
}

var to, from string

func init() {
	sendCmd.Flags().StringVarP(&from, "from", "f", "jumposhi", "address to take coins")
	sendCmd.Flags().StringVarP(&to, "to", "t", "jumposhi", "address to send coins")

	RootCmd.AddCommand(sendCmd)
}
