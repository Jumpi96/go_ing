package internal

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var blocksBucket string = "blocks"
var db *bolt.DB

type Repository interface {
	SaveNewBlock(*Block) error
	GetLastBlockHash() ([]byte, error)
	GetBlock([]byte) (*Block, error)
}

type BoltDBRepository struct{}

func InitDB() *BoltDBRepository {
	db, _ = bolt.Open("blockychain.db", 0600, nil)
	return &BoltDBRepository{}
}

func (r *BoltDBRepository) SaveNewBlock(block *Block) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			b, _ = tx.CreateBucket([]byte(blocksBucket))
		}
		serialized, err := block.serializeBlock()
		if err != nil {
			log.Panic(fmt.Sprintf("Serializing error: %v", err))
		}
		err = b.Put([]byte(block.Hash), serialized)
		err = b.Put([]byte("l"), []byte(block.Hash))
		return err
	})
	return err
}

func (r *BoltDBRepository) GetLastBlockHash() ([]byte, error) {
	var blockHash []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b != nil {
			blockHash = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return blockHash, nil
}

func (r *BoltDBRepository) GetBlock(hash []byte) (*Block, error) {
	var block []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		block = b.Get(hash)
		return nil
	})

	if err != nil {
		return nil, err
	}

	deserialized, err := deserializeBlock(block)

	if err != nil {
		log.Panic(fmt.Sprintf("Deserialized error: %v", err))
	}

	return deserialized, nil
}
