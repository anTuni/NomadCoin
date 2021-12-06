package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/anTuni/NomadCoin/db"
	"github.com/anTuni/NomadCoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(b)
	utils.HandleErr(err)
}
func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(
			func() {
				b = &blockchain{"", 0}
				checkpoint := db.Checkpoint()
				if checkpoint == nil {
					b.AddBlock("Genesis Block")
				} else {
					fmt.Println("Restoring...")
					b.restore(checkpoint)
				}
			})
	}
	fmt.Printf("Newest Hash:%s \nHeight: %d\n", b.NewestHash, b.Height)

	return b
}
