package main

import (
	"fmt"

	internal "./internal"
)

var r = &internal.BoltDBRepository{}

func main() {
	blockchain := internal.NewBlockchain(r, 1)
	blockchain.AddBlock(r, internal.NewBlock([]byte(""), "06/27/2020", internal.Data{Value: "This is a second block"}))
	blockchain.AddBlock(r, internal.NewBlock([]byte(""), "06/27/2020", internal.Data{Value: "This is the third block"}))

	iterator := blockchain.Iterator()
	count := 1
	for iterator.Next(r) != nil {
		count++
	}

	fmt.Println(count)
	fmt.Println(blockchain.IsChainValid(r))

}
