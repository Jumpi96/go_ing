package internal

import (
	"bytes"
	"time"
)

type Blockchain struct {
	Tip        []byte
	difficulty int
}

var genesisPreviousHash []byte = []byte("0")

func NewBlockchain(r Repository, difficulty int) *Blockchain {
	tip, _ := r.GetLastBlockHash()

	if bytes.Equal(tip, []byte("")) {
		genesis := createGenesisBlock()
		tip = genesis.Hash
		r.SaveNewBlock(genesis)
	}

	return &Blockchain{tip, difficulty}
}

func createGenesisBlock() *Block {
	genesis := &Block{
		PreviousHash: genesisPreviousHash,
		Timestamp:    time.Now().Format("02/01/2006"),
		Data:         Data{Value: "This is a genesis block."},
		Nonce:        0,
	}
	genesis.Hash = calculateHash(genesis)
	return genesis
}

func (c Blockchain) GetLatestBlock(r Repository) *Block {
	block, _ := r.GetBlock(c.Tip)
	return block
}

func (c *Blockchain) AddBlock(r Repository, newBlock *Block) error {
	newBlock.PreviousHash = c.Tip
	newBlock.mineBlock(c.difficulty)
	err := r.SaveNewBlock(newBlock)
	c.Tip = newBlock.Hash

	return err
}

type BlockchainIterator struct {
	currentHash []byte
	repository  Repository
}

func (c *Blockchain) Iterator(r Repository) *BlockchainIterator {
	return &BlockchainIterator{c.Tip, r}
}

func (i *BlockchainIterator) Next() *Block {
	if !bytes.Equal(i.currentHash, genesisPreviousHash) {
		block, _ := i.repository.GetBlock(i.currentHash)
		i.currentHash = block.PreviousHash
		return block
	}
	return nil
}

func (c *Blockchain) IsChainValid(r Repository) bool {
	var previousBlock *Block
	iterator := c.Iterator(r)
	currentBlock := iterator.Next()

	for currentBlock != nil {
		if !bytes.Equal(currentBlock.Hash, calculateHash(currentBlock)) {
			return false
		}

		previousBlock = iterator.Next()
		if previousBlock != nil {
			if !bytes.Equal(currentBlock.PreviousHash, previousBlock.Hash) {
				return false
			}
			currentBlock = previousBlock
		} else {
			break
		}
	}
	return true
}
