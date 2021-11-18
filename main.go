package main

import (
	"fmt"

	"github.com/anTuni/NomadCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("third Block")
	chain.AddBlock("forth Block")
	for _, block := range chain.AllBolocks() {
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %s \n", block.Hash)
		fmt.Printf("PrevHash : %s \n", block.PrevHash)
	}
}
