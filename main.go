package main

import (
	"fmt"

	"github.com/k1nky/simplechain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.Append("1st transaction")
	bc.Append("2nd transaction")
	fmt.Printf("%s", bc)
}
