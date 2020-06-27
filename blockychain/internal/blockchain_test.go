package internal

import (
	"testing"
)

func TestAddBlocksToBlockchain(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is a second block"}))
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is the third block"}))

	expectedValue := 3
	if len(blockchain.chain) != expectedValue {
		t.Errorf("Number of blocks was incorrect, got: %v, want: %v.", len(blockchain.chain), expectedValue)
	}

	thirdBlockValue := "This is the third block"
	if blockchain.getLatestBlock().data.value != thirdBlockValue {
		t.Errorf("Number of blocks was incorrect, got: %v, want: %v.", thirdBlockValue, expectedValue)
	}
}

func TestIntegrityBeforeAndAfterTamperingIt(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is a second block"}))
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is the third block"}))

	if !blockchain.IsChainValid() {
		t.Errorf("Blockchain is not valid.")
	}

	blockchain.chain[1].data = data{value: "This is the genesis block"}
	blockchain.chain[1].hash = calculateHash(blockchain.chain[1])

	if blockchain.IsChainValid() {
		t.Errorf("Blockchain is valid but it should not be because it was tampered.")
	}
}
