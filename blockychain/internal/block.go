package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"strings"
)

type Data struct {
	Value string
}

type Block struct {
	PreviousHash []byte
	Timestamp    string
	Data         Data
	Hash         []byte
	Nonce        int
}

func NewBlock(timestamp string, data Data) *Block {
	block := &Block{
		Timestamp: timestamp,
		Data:      data,
		Nonce:     0,
	}
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(b *Block) []byte {
	input := fmt.Sprintf("%v%v%v%v", b.PreviousHash, b.Timestamp, b.Data, b.Nonce)
	return SHA256(input)
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

func (b *Block) serializeBlock() []byte {
	buffer := bytes.Buffer{}
	e := gob.NewEncoder(&buffer)
	err := e.Encode(b)
	if err != nil {
		panic(fmt.Sprintf("Serializing %v", err))
	}
	return buffer.Bytes()
}

func deserializeBlock(d []byte) *Block {
	b := &Block{}
	buffer := bytes.Buffer{}
	buffer.Write(d)
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(&b)
	if err != nil {
		panic(fmt.Sprintf("Deserializing %v %v", d, err))
	}
	return b
}
