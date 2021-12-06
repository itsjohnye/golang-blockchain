// a simple demo to illustrate the basic concept of blockchain
package main

import (
	"fmt"
	"github/itsjohnye/golang-blickchain/2-proof-of-work/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
