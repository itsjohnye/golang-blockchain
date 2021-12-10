//proof of work algorithm
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)


const Difficulty = 12 

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) 

	pow := &ProofOfWork{b, target} 

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreviousHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 { // math.MaxInt64 equals to 2^63 - 1
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash) // "\r" means return, it is used to move the cursor to the beginning of the line, so that we can see the progress of finding the appropriate hash, AKA "mining".

		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 { // if the hash is smaller than the target, then we have a valid hash
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	fmt.Printf("nonce=%v\n", nonce)

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}
