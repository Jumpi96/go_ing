package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Transaction should have at least one TXInput and one TXOutput
type Transaction struct {
	ID  []byte
	In  []TXInput
	Out []TXOutput
}

// TXInput references a previous transaction and provides data for the TXOutput.ScriptKey
type TXInput struct {
	TxID            []byte
	OutIndex        int    // Index of the specific TXOutput in the transaction
	ScriptSignature string // Data to be used by ScriptKey
}

// TXOutput are where coins are stored
type TXOutput struct {
	Value     int
	ScriptKey string // A puzzle to unlock the output in Script language
}

var subsidy int = 50

// SetID sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func newCoinbaseTransaction(to string) *Transaction {
	data := fmt.Sprintf("Reward to '%s'", to)
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := &Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return tx
}

func (in *TXInput) outputUnlockableWith(unlockingData string) bool {
	return in.ScriptSignature == unlockingData
}

func (out *TXOutput) unlockableWith(unlockingData string) bool {
	return out.ScriptKey == unlockingData
}

func NewUTXOTransaction(r Repository, bc *Blockchain, from, to string, amount int) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.findSpendableOutputs(r, from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, _ := hex.DecodeString(txid)

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
