package internal

import (
	"testing"
)

func TestMineBlock(t *testing.T) {
	difficulty := 1
	b := NewBlock("", "06/27/2020", data{value: "This is a first block"})
	b.mineBlock(difficulty)
	if b.nonce == 0 {
		t.Errorf("I don't think mining is working.")
	}
}
