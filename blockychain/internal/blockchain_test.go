package internal

import (
	"testing"
)

func generateSimpleBlockchain() *Blockchain {
	blockchain := NewBlockchain(1)
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is a second block"}))
	blockchain = blockchain.AddBlock(NewBlock("", "06/27/2020", data{value: "This is the third block"}))
	return blockchain
}

func TestAddBlocksToBlockchain(t *testing.T) {
	blockchain := generateSimpleBlockchain()

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
	blockchain := generateSimpleBlockchain()

	if !blockchain.IsChainValid() {
		t.Errorf("Blockchain is not valid.")
	}

	blockchain.chain[1].data = data{value: "This is the genesis block"}
	blockchain.chain[1].hash = calculateHash(blockchain.chain[1])

	if blockchain.IsChainValid() {
		t.Errorf("Blockchain is valid but it should not be because it was tampered.")
	}
}

func TestIntegrityAfterBreakingIt(t *testing.T) {
	blockchain := generateSimpleBlockchain()

	blockchain.chain[2].previousHash = "a"

	if blockchain.IsChainValid() {
		t.Errorf("Blockchain is valid but it should not be because it was broken.")
	}
}

func TestIntegrityAfterReplacingGenesisBlock(t *testing.T) {
	blockchain := generateSimpleBlockchain()

	blockchain.chain[0] = NewBlock(blockchain.chain[2].hash, "06/27/2020", data{value: "I'm your genesis block."})

	if blockchain.IsChainValid() {
		t.Errorf("Blockchain is valid but it should not be because its genesis block was replaced.")
	}
}
