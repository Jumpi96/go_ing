package internal

import (
	"testing"
)

func TestMineBlock(t *testing.T) {
	difficulty := 1
	b := NewBlock([]byte(""), "06/27/2020", Data{Value: "This is a first block"})
	b.mineBlock(difficulty)
	if b.Nonce == 0 {
		t.Errorf("I don't think mining is working.")
	}
}
