package internal

import (
	"crypto/sha256"
	"fmt"
)

type data struct {
	value string
}

type Block struct {
	previousHash string
	timestamp    string
	data         data
	hash         string
}

func NewBlock(previous string, timestamp string, data data) *Block {
	block := &Block{
		previousHash: previous,
		timestamp:    timestamp,
		data:         data,
	}
	block.hash = calculateHash(block)
	return block
}

func calculateHash(b *Block) string {
	h := sha256.New()
	input := fmt.Sprintf("%v%v%v", b.previousHash, b.timestamp, b.data)
	return fmt.Sprintf("%x", h.Sum([]byte(input)))
}
