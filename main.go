package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	base16Hash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = base16Hash
	fmt.Println(genesisBlock)
}
