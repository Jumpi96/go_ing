package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type data struct {
	value string
}

type Block struct {
	previousHash string
	timestamp    string
	data         data
	hash         string
	nonce        int
}

func NewBlock(previous string, timestamp string, data data) *Block {
	block := &Block{
		previousHash: previous,
		timestamp:    timestamp,
		data:         data,
		nonce:        0,
	}
	block.hash = calculateHash(block)
	return block
}

func calculateHash(b *Block) string {
	input := fmt.Sprintf("%v%v%v%v", b.previousHash, b.timestamp, b.data, b.nonce)
	return SHA256(input)
}

func SHA256(text string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func (b *Block) mineBlock(difficulty int) {
	goal := strings.Repeat("0", difficulty)
	for b.hash[:difficulty] != goal {
		b.nonce++
		b.hash = calculateHash(b)
	}
}
