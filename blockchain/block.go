package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/anTuni/NomadCoin/db"
	"github.com/anTuni/NomadCoin/utils"
)

const difficulty = 3

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
	Timestamp  int    `json:"timestamp"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		hash := utils.Hash(b)
		b.Timestamp = int(time.Now().Unix())
		fmt.Printf("Target : %s\nblockAsString : %s , Hash : %s\nNonce : %d", hash, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: Blockchain().difficulty(),
		Nonce:      0,
	}
	block.mine()
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
