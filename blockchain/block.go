package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

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
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("blockAsString : %s\nHash : %s\nTarget : %s\nNonce : %d", blockAsString, hash, target, b.Nonce)
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
		Difficulty: difficulty,
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
