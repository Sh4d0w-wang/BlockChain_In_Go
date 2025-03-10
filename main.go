package main

import (
	"BlockChain_In_Go/consensus"
	"fmt"
)

func main() {
	blockchain := consensus.NewBlockChain()
	blockchain.AddBlock([]byte("Send 1 BTC To WZM"))
	blockchain.AddBlock([]byte("Send 2 BTC To MZW"))
	for i, block := range blockchain.Blocks {
		fmt.Printf("--------------------------------%v--------------------------------\n", i)
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		pow := consensus.NewPoofOfWork(block)
		fmt.Printf("Pow Validate:%v\n", pow.Validate())
		fmt.Printf("--------------------------------END--------------------------------\n")
	}
}
