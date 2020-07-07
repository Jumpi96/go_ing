package cmd

import (
	internal "../internal"

	"github.com/spf13/cobra"
)

var blockchain *internal.Blockchain
var repository = &internal.BoltDBRepository{}
var difficulty = 1

var RootCmd = &cobra.Command{
	Use:   "blockychain",
	Short: "Blockychain is a blockchain, for real",
}

func Init(r *internal.BoltDBRepository) {
	blockchain = internal.NewBlockchain(r, difficulty)
	RootCmd.Execute()
}
