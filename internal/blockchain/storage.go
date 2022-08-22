package blockchain

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

const (
	BlocksBucket = "blocks"
)

type Storage interface {
	PutBlock(block *Block)
	Open(connString string) error
	Tail() *Block
	// Next(current *Block) *Block
}

type SliceStorage struct {
	blocks []*Block
}

type PeristentStorage struct {
	db *bolt.DB
}

func (s *SliceStorage) Open(connString string) error {
	return nil
}

func (s *SliceStorage) PutBlock(block *Block) {
	s.blocks = append(s.blocks, block)
}

func (s *SliceStorage) Tail() *Block {
	if len(s.blocks) == 0 {
		return nil
	}
	return s.blocks[len(s.blocks)-1]
}

func (s *PeristentStorage) Open(connString string) error {
	var err error
	if s.db, err = bolt.Open(connString, 0600, nil); err != nil {
		return err
	}
	err = s.db.Update(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(BlocksBucket)); b == nil {
			b, err = tx.CreateBucket([]byte(BlocksBucket))
		}
		return nil
	})

	return err
}

func (s *PeristentStorage) PutBlock(block *Block) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		if err := b.Put(block.Hash, block.Serialize()); err != nil {
			return err
		}
		if err := b.Put([]byte("tail"), block.Hash); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Print(err)
	}
}

func (s *PeristentStorage) Tail() *Block {
	var block *Block

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		if b == nil {
			block = nil
		} else {
			block = DeserializeBlock(b.Get([]byte("tail")))
		}
		return nil
	})
	if err != nil {
		fmt.Print(err)
		block = nil
	}

	return block
}
