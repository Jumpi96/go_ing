package cmd

import (
	internal "../internal"

	"github.com/spf13/cobra"
)

var repository = &internal.BoltDBRepository{}
var difficulty = 1

var RootCmd = &cobra.Command{
	Use:   "blockychain",
	Short: "Blockychain is a blockchain, for real",
}

func Init(r *internal.BoltDBRepository) {
	RootCmd.Execute()
}
