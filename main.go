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

type blockchain struct {
	blocks []block
}

func (b *blockchain) getList() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s \n", block.data)
		fmt.Printf("Hash: %s \n", block.hash)
		fmt.Printf("Prev Hash: %s \n", block.prevHash)
	}
}
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}
func (b *blockchain) getNewBlcok(data string) block {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(data + b.getLastHash()))
	newBlock.hash = fmt.Sprintf("%x", hash)
	return newBlock
}

func (b *blockchain) addBlock(data string) {
	b.blocks = append(b.blocks, b.getNewBlcok(data))
}

func (b *block) assignHash(data, prevHash string) {
}
func main() {
	firstBlock := blockchain{}
	firstBlock.addBlock("tfirst")
	firstBlock.addBlock("second")
	firstBlock.addBlock("third")
	firstBlock.getList()
}
