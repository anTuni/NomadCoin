package blockchain

import (
	"errors"
	"time"

	"github.com/anTuni/NomadCoin/utils"
)

const (
	rewardForMiner int = 50
)

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

func makeCoinbaseTx(address string) *Tx {
	TxIns := []*TxIn{
		{"COINBASE", rewardForMiner},
	}
	TxOuts := []*TxOut{
		{address, rewardForMiner},
	}
	t := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     TxIns,
		TxOuts:    TxOuts,
	}
	t.getId()
	return &t
}
func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}

	return nil, nil
}
func (m *mempool) AddTx(to string, amount int) error {
	Tx, err := makeTx("taeyun", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, Tx)
	return nil
}
