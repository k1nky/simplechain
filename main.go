package main

import (
	"fmt"

	"github.com/k1nky/simplechain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.Append("1st transaction")
	bc.Append("2nd transaction")
	// fmt.Printf("%+v", bc)
	it := bc.Blocks.Iterator()
	for {
		block := it.Next()
		if block.IsEmpty() {
			break
		}
		fmt.Printf("%s\n", block)
	}
}
