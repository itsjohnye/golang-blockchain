package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

//Genesis Block is the first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
