package internal

import (
	"testing"
	"time"
)

func TestMineBlock(t *testing.T) {
	difficulty := 1
	b := NewBlock("06/27/2020", []*Transaction{})
	b.mineBlock(difficulty)
	if b.Nonce == 0 {
		t.Errorf("I don't think mining is working.")
	}
}

func TestBlockSerialization(t *testing.T) {
	block := NewBlock(time.Now().Format("02/01/2006"), []*Transaction{})

	serialized, err := block.serializeBlock()

	if serialized == nil {
		t.Errorf("Serialization isn't working: %v.", err)
	}

	deserialized, err := deserializeBlock(serialized)

	if deserialized == nil {
		t.Errorf("Deserialization isn't working: %v.", err)
	}

	serialized[1] = serialized[0]
	notDeseralizable, err := deserializeBlock(serialized)

	if notDeseralizable != nil {
		t.Errorf("Deserialization shouldn't work with bad data: %v.", err)
	}
}
