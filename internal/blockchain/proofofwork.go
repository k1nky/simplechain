package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const (
	TargetBits = 16
	MaxNonce   = math.MaxInt64
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBits))
	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

func (pow *ProofOfWork) prepare(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.block.ParentHash,
			pow.block.Data,
			IntToBytes(pow.block.Timestamp),
			IntToBytes(int64(TargetBits)),
			IntToBytes(int64(nonce)),
		},
		[]byte{},
	)
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < MaxNonce {
		data := pow.prepare(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("Found the hash %x with nonce %d\n", hash, nonce)
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepare(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
