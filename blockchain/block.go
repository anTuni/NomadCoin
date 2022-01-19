package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/anTuni/NomadCoin/db"
	"github.com/anTuni/NomadCoin/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

func persistBlock(b *Block) {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := utils.Hash(b)
		b.Timestamp = int(time.Now().Unix())
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height, diff int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: diff,
		Nonce:      0,
	}
	block.mine()
	block.Transactions = Mempool.TxsToConfirm()
	persistBlock(block)
	return block
}

var ErrorNotFound error = errors.New("block not found")

func FindBlock(hash string) (*Block, error) {
	data := db.Block(hash)
	if data == nil {
		return nil, ErrorNotFound
	}
	block := &Block{}
	utils.FromBytes(block, data)
	return block, nil
}
