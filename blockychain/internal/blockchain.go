package internal

import (
	"reflect"
)

type Blockchain struct {
	chain      []*Block
	difficulty int
}

func NewBlockchain(difficulty int) *Blockchain {
	return &Blockchain{chain: []*Block{createGenesisBlock()}, difficulty: difficulty}
}

func createGenesisBlock() *Block {
	return NewBlock("0", "06/27/2020", data{value: "This is a genesis block."})
}

func (c Blockchain) getLatestBlock() *Block {
	return c.chain[len(c.chain)-1]
}

func (c *Blockchain) AddBlock(newBlock *Block) *Blockchain {
	newBlock.previousHash = c.getLatestBlock().hash
	newBlock.mineBlock(c.difficulty)
	c.chain = append(c.chain, newBlock)
	return c
}

func (c *Blockchain) IsChainValid() bool {
	for i, currentBlock := range c.chain {
		if i != 0 {
			previousBlock := c.chain[i-1]

			if currentBlock.hash != calculateHash(currentBlock) {
				return false
			}

			if currentBlock.previousHash != previousBlock.hash {
				return false
			}

		} else {
			if !(reflect.DeepEqual(currentBlock, createGenesisBlock())) {
				return false
			}
		}
	}
	return true
}
