package internal

import (
	"bytes"
	"encoding/hex"
	"time"
)

type Blockchain struct {
	Tip        []byte
	difficulty int
}

var genesisPreviousHash = []byte{}

func NewBlockchain(r Repository, address string, difficulty int) *Blockchain {
	genesis := createGenesisBlock(address)
	genesis.mineBlock(difficulty)
	r.SaveNewBlock(genesis)
	return &Blockchain{genesis.Hash, difficulty}
}

func GetBlockchain(r Repository, difficulty int) *Blockchain {
	tip, _ := r.GetLastBlockHash()
	if bytes.Equal(tip, []byte("")) {
		return nil
	}
	return &Blockchain{tip, difficulty}
}

func createGenesisBlock(address string) *Block {
	genesis := &Block{
		PreviousHash: genesisPreviousHash,
		Timestamp:    time.Now().Format("02/01/2006"),
		Transactions: []*Transaction{newCoinbaseTransaction(address)},
		Nonce:        0,
	}
	genesis.Hash = calculateHash(genesis)
	return genesis
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

// IsChainValid checks if blockchain PoW works
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

func (tx Transaction) isCoinbase() bool {
	return len(tx.In) == 1 && len(tx.In[0].TxID) == 0 && tx.In[0].OutIndex == -1
}

func (c *Blockchain) GetBalance(r Repository, address string) int {
	balance := 0
	UTXOs := c.findUTXO(r, address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}

func (c *Blockchain) findUTXO(r Repository, address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := c.getUnspentTransactions(r, address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Out {
			if out.unlockableWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (c *Blockchain) getUnspentTransactions(r Repository, address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	iterator := c.Iterator(r)

	for {
		block := iterator.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

			for outIdx, out := range tx.Out {
				if spentTXOs[txID] != nil { // Was the output spent?
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							break
						}
					}
				}

				if out.unlockableWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.isCoinbase() == false {
				for _, in := range tx.In {
					if in.outputUnlockableWith(address) {
						inTxID := hex.EncodeToString(in.TxID)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.OutIndex)
					}
				}
			}
		}

		if len(block.PreviousHash) == 0 {
			break
		}
	}

	return unspentTXs
}

func (c *Blockchain) findSpendableOutputs(r Repository, address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := c.getUnspentTransactions(r, address)
	accumulated := 0

	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Out {
			if out.unlockableWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break
				}
			}
		}
	}

	return accumulated, unspentOutputs
}
