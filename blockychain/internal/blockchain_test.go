package internal

import (
	"testing"
)

var r = &BoltDBRepository{}

func generateSimpleBlockchain() *Blockchain {
	blockchain := NewBlockchain(r, 1)
	blockchain.AddBlock(r, NewBlock([]byte(""), "06/27/2020", Data{Value: "This is a second block"}))
	blockchain.AddBlock(r, NewBlock([]byte(""), "06/27/2020", Data{Value: "This is the third block"}))
	return blockchain
}

func TestAddBlocksToBlockchain(t *testing.T) {
	blockchain := generateSimpleBlockchain()

	expectedValue := 3

	iterator := BlockchainIterator{blockchain.Tip}
	count := 1
	for iterator.Next(r) != nil {
		count++
	}

	if count != expectedValue {
		t.Errorf("Number of blocks was incorrect, got: %v, want: %v.", count, expectedValue)
	}

	thirdBlockValue := "This is the third block"
	if blockchain.GetLatestBlock(r).Data.Value != thirdBlockValue {
		t.Errorf("Number of blocks was incorrect, got: %v, want: %v.", thirdBlockValue, expectedValue)
	}
}
