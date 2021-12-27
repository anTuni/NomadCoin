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

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
}

func recalDifficulty(b *blockchain) int {
	blocks := Blocks(b)
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

func persistBlockchain(b *blockchain) {
	db.SaveBlockchain(utils.ToBytes(b))
}

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyinterval == 0 {
		return recalDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func Blocks(b *blockchain) []*Block {
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
func Txs(b *blockchain) []*Tx {
	var Txs []*Tx
	for _, block := range Blocks(b) {
		Txs = append(Txs, block.Transactions...)
	}
	return Txs
}

func FindTx(b *blockchain, targetID string) *Tx {
	Txs := Txs(b)
	for _, Tx := range Txs {
		if Tx.Id == targetID {
			return Tx
		}
	}
	return nil
}
func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var UTxOuts []*UTxOut
	createdTxIds := make(map[string]bool)
	for _, block := range Blocks(Blockchain()) {
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

func BalanceByAddress(address string, b *blockchain) int {
	var amount int
	TxOuts := UTxOutsByAddress(address, b)
	for _, TxOut := range TxOuts {
		amount += TxOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
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
	return b
}
