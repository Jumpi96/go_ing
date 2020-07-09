package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	PreviousHash []byte
	Timestamp    string
	Transactions []*Transaction
	Hash         []byte
	Nonce        int
}

func NewBlock(timestamp string, txs []*Transaction) *Block {
	block := &Block{
		Timestamp:    timestamp,
		Transactions: txs,
		Nonce:        0,
	}
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(b *Block) []byte {
	input := fmt.Sprintf("%v%v%x%v", b.PreviousHash, b.Timestamp, b.hashTransactions(), b.Nonce)
	return SHA256(input)
}

func (b *Block) hashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func SHA256(text string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(text))
	return algorithm.Sum(nil)
}

func (b *Block) mineBlock(difficulty int) {
	goal := strings.Repeat("0", difficulty)
	for hex.EncodeToString(b.Hash)[:difficulty] != goal {
		b.Nonce++
		b.Hash = calculateHash(b)
	}
}

func (b *Block) serializeBlock() ([]byte, error) {
	buffer := bytes.Buffer{}
	e := gob.NewEncoder(&buffer)
	err := e.Encode(b)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func deserializeBlock(d []byte) (*Block, error) {
	b := &Block{}
	buffer := bytes.Buffer{}
	buffer.Write(d)
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(&b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
