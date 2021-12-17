package blockchain

import (
	"errors"
	"fmt"
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

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := utils.Hash(b)
		b.Timestamp = int(time.Now().Unix())
		fmt.Printf("Target : %s\nblockAsString : %s , Hash : %s\nNonce : %d\n", hash, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: Blockchain().difficulty(),
		Nonce:      0,
	}
	block.mine()
	block.Transactions = Mempool.TxsToConfirm()
	block.persist()
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
