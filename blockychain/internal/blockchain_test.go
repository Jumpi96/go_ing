package internal

import (
	"bytes"
	"testing"
	"time"
)

type mockRepository struct{}

//var saveNewBlockMock func(block *Block) error

func (m *mockRepository) SaveNewBlock(block *Block) error {
	return nil //saveNewBlockMock(block)
}

var getLastBlockHashMock func() ([]byte, error)

func (m *mockRepository) GetLastBlockHash() ([]byte, error) {
	return getLastBlockHashMock()
}

var getBlockMock func([]byte) (*Block, error)

func (m *mockRepository) GetBlock(hash []byte) (*Block, error) {
	return getBlockMock(hash)
}

func TestNewBlockchain(t *testing.T) {
	r := &mockRepository{}
	difficulty := 1
	address := "testing"
	genesis := createGenesisBlock(address)

	blockchain := NewBlockchain(r, address, difficulty)

	if !bytes.Equal(blockchain.Tip, genesis.Hash) {
		t.Errorf("Blockchain tip is incorrect, got: %v, want: %v.", blockchain.Tip, genesis.Hash)
	}

	if blockchain.difficulty != difficulty {
		t.Errorf("Blockchain difficulty is incorrect, got: %v, want: %v.", blockchain.difficulty, difficulty)
	}
}

func TestGetBlockchainThatExistsAndThatNotExists(t *testing.T) {
	r := &mockRepository{}
	difficulty := 1
	address := "testing"

	blockchain := NewBlockchain(r, address, difficulty)

	getLastBlockHashMock = func() ([]byte, error) {
		// Genesis block hash with "testing" address
		genesis := createGenesisBlock(address)
		return genesis.Hash, nil
	}

	blockchainGot := GetBlockchain(r, difficulty)

	if !bytes.Equal(blockchain.Tip, blockchainGot.Tip) {
		t.Errorf("Blockchain tip is incorrect, got: %v, want: %v.", blockchain.Tip, blockchainGot.Tip)
	}

	getLastBlockHashMock = func() ([]byte, error) { return nil, nil }

	blockchainGot = GetBlockchain(r, difficulty)

	if blockchainGot != nil {
		t.Errorf("Blockchain shouldn't exist.")
	}
}

func TestAddBlockToBlockchain(t *testing.T) {
	r := &mockRepository{}

	blockchain := NewBlockchain(r, "testing", 1)
	previousTip := blockchain.Tip

	// TODO: Complete with transactions
	newBlock := NewBlock(time.Now().Format("02/01/2006"), []*Transaction{})
	err := blockchain.AddBlock(r, newBlock)

	if err != nil {
		t.Errorf("Block couldn't be added: %e.", err)
	}

	if !bytes.Equal(blockchain.Tip, newBlock.Hash) {
		t.Errorf("Blockchain tip is incorrect, got: %v, want: %v.", blockchain.Tip, newBlock.Hash)
	}

	if newBlock.Nonce == 0 {
		t.Errorf("Block wasn't mined!")
	}

	if !bytes.Equal(previousTip, newBlock.PreviousHash) {
		t.Errorf("New block previous hash is incorrect, got: %v, want: %v.", previousTip, newBlock.PreviousHash)
	}
}

func TestIterator(t *testing.T) {
	r := &mockRepository{}

	blockchain := NewBlockchain(r, "testing", 1)
	iterator := blockchain.Iterator(r)

	getBlockMock = func([]byte) (*Block, error) {
		return createGenesisBlock("testing"), nil
	}

	aBlock := iterator.Next()

	if aBlock == nil {
		t.Errorf("Iterator should return the genesis block.")
	}

	noBlock := iterator.Next()

	if noBlock != nil {
		t.Errorf("Iterator should not return more blocks.")
	}

}
func TestIsChainValid(t *testing.T) {
	r := &mockRepository{}

	blockchain := NewBlockchain(r, "testing", 1)

	getBlockMock = func([]byte) (*Block, error) {
		return createGenesisBlock("testing"), nil
	}

	if !blockchain.IsChainValid(r) {
		t.Errorf("Blockchain PoW is not valid!")
	}

	getBlockMock = func([]byte) (*Block, error) {
		genesis := createGenesisBlock("testing")
		genesis.Timestamp = "Not date"
		return genesis, nil
	}

	if blockchain.IsChainValid(r) {
		t.Errorf("A block was tampered!")
	}

	// TODO: test a broken blockchain.
}
