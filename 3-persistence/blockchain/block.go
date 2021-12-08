package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)



type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//Genesis Block is the first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize the block into bytes in order to be stored by using BadgerDB
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
