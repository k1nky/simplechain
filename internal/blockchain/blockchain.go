package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Data       []byte
	Hash       []byte
	Nonce      int
	ParentHash []byte
	Timestamp  int64
}

type Blockchain struct {
	Blocks Storage
}

func (b *Block) String() string {
	return fmt.Sprintf("%s [%x/%d]", b.Data, b.Hash, b.Nonce)
}

func (b *Block) IsEmpty() bool {
	return len(b.Hash) == 0
}

func (b *Block) Serialize() []byte {
	result := bytes.Buffer{}
	gob.NewEncoder(&result).Encode(b)

	return result.Bytes()
}

func DeserializeBlock(raw []byte) *Block {
	block := &Block{}
	gob.NewDecoder(bytes.NewReader(raw)).Decode(block)
	return block
}

func NewBlock(data string, parentHash []byte) (block *Block) {
	block = &Block{
		Data:       []byte(data),
		Hash:       []byte{},
		ParentHash: parentHash,
		Timestamp:  time.Now().Unix(),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return
}

func (bc *Blockchain) Append(data string) {
	tail := bc.Blocks.Tail()
	newBlock := NewBlock(data, tail.Hash)
	bc.Blocks.PutBlock(newBlock)
}

// func (bc *Blockchain) String() string {
// 	s := ""
// 	for _, v := range bc.Blocks {
// 		s += fmt.Sprintln(v)
// 	}
// 	return s
// }

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		Blocks: &PeristentStorage{},
	}
	bc.Blocks.Open("local.db")
	if tail := bc.Blocks.Tail(); tail.IsEmpty() {
		bc.Blocks.PutBlock(NewGenesisBlock())
	}

	return bc
}
