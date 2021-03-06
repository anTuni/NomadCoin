package blockchain

import (
	"encoding/json"
	"net/http"
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
	m                 sync.Mutex
}

var b *blockchain
var once sync.Once

func Blockchain() *blockchain {
	once.Do(
		func() {
			b = &blockchain{Height: 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				b.restore(checkpoint)
			}
		})
	return b
}

func (b *blockchain) AddBlock() *Block {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
	return block
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
func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
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
func Txs(b *blockchain) []*Tx {
	var Txs []*Tx
	for _, block := range Blocks(b) {
		Txs = append(Txs, block.Transactions...)
	}
	return Txs
}

func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()
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
func BalanceByAddress(address string, b *blockchain) int {
	var amount int
	TxOuts := UTxOutsByAddress(address, b)
	for _, TxOut := range TxOuts {
		amount += TxOut.Amount
	}
	return amount
}
func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var UTxOuts []*UTxOut
	createdTxIds := make(map[string]bool)
	for _, block := range Blocks(Blockchain()) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					continue
				}
				if FindTx(b, input.TxId).TxOuts[input.Index].Address == address {
					createdTxIds[input.TxId] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Address == address {
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

func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()
	utils.HandleErr(json.NewEncoder(rw).Encode(b))
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()
	b.CurrentDifficulty = newBlocks[0].Difficulty
	b.NewestHash = newBlocks[0].Hash
	b.Height = len(newBlocks)
	persistBlockchain(b)

	db.EmptyBlocks()

	for _, block := range newBlocks {
		persistBlock(block)
	}
}
func (b *blockchain) AddPeerBlock(newBlock *Block) {
	Mempool().m.Lock()
	defer Mempool().m.Unlock()
	b.m.Lock()
	defer b.m.Unlock()
	b.Height += 1
	b.CurrentDifficulty = newBlock.Difficulty
	b.NewestHash = newBlock.Hash
	persistBlockchain(b)
	persistBlock(newBlock)

	for _, Tx := range newBlock.Transactions {
		_, ok := Mempool().Txs[Tx.Id]
		if ok {
			delete(Mempool().Txs, Tx.Id)
		}
	}
}

func persistBlockchain(b *blockchain) {
	db.SaveBlockchain(utils.ToBytes(b))
}
