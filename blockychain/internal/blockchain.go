package internal

import (
	"bytes"
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
	return NewBlock(genesisPreviousHash, "06/27/2020", Data{Value: "This is a genesis block."})
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
}

func (c *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{c.Tip}
}

func (i *BlockchainIterator) Next(r Repository) *Block {
	block, _ := r.GetBlock(i.currentHash)
	if !bytes.Equal(block.PreviousHash, genesisPreviousHash) {
		i.currentHash = block.PreviousHash
		return block
	}
	return nil
}

func (c *Blockchain) IsChainValid(r Repository) bool {
	var previousBlock *Block
	iterator := c.Iterator()
	currentBlock := iterator.Next(r)

	for currentBlock != nil {
		if !bytes.Equal(currentBlock.Hash, calculateHash(currentBlock)) {
			return false
		}

		previousBlock = iterator.Next(r)
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
