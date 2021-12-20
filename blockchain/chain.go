package blockchain

import (
	"fmt"
	"sync"

	"github.com/anTuni/NomadCoin/db"
	"github.com/anTuni/NomadCoin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyinterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) recalDifficulty() int {
	blocks := b.Blocks()
	newestBlock := blocks[0]
	lastCalculatedBlock := blocks[difficultyinterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastCalculatedBlock.Timestamp / 60)
	expectedTime := blockInterval * difficultyinterval

	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyinterval == 0 {
		return b.recalDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}
func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}
func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}
func (b *blockchain) UTxOutsByAddress(address string) []*UTxOut {
	var UTxOuts []*UTxOut
	createdTxIds := make(map[string]bool)
	for _, block := range Blockchain().Blocks() {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Owner == address {
					createdTxIds[input.TxId] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Owner == address {
					if _, ok := createdTxIds[tx.Id]; !ok {
						UTxOut := &UTxOut{TxId: tx.Id, Index: index, Amount: output.Amount}
						if !isOnMempool(UTxOut) {
							UTxOuts = append(UTxOuts, UTxOut)
						}
					}
				}
			}
		}
	}

	return UTxOuts
}
func (b *blockchain) BalanceByAddress(address string) int {
	var amount int
	TxOuts := b.UTxOutsByAddress(address)
	for _, TxOut := range TxOuts {
		amount += TxOut.Amount
	}
	return amount
}
func Blockchain() *blockchain {
	if b == nil {
		once.Do(
			func() {
				b = &blockchain{Height: 0}
				checkpoint := db.Checkpoint()
				if checkpoint == nil {
					b.AddBlock()
				} else {
					fmt.Println("Restoring...")
					b.restore(checkpoint)
				}
			})
	}
	fmt.Printf("Newest Hash:%s \nHeight: %d\n", b.NewestHash, b.Height)

	return b
}
