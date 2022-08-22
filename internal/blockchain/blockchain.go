package blockchain

import (
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
	Blocks []*Block
}

func (b *Block) String() string {
	return fmt.Sprintf("%s [%x/%d]", b.Data, b.Hash, b.Nonce)
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
	parentHash := bc.Blocks[len(bc.Blocks)-1]
	bc.Blocks = append(bc.Blocks, NewBlock(data, parentHash.Hash))
}

func (bc *Blockchain) String() string {
	s := ""
	for _, v := range bc.Blocks {
		s += fmt.Sprintln(v)
	}
	return s
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}
