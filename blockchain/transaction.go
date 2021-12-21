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
	TxId  string `json:"txid"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}
type UTxOut struct {
	TxId   string `json:"txid"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

func isOnMempool(UTxOut *UTxOut) bool {
	exists := false
Outer:
	for _, Tx := range Mempool.Txs {
		for _, input := range Tx.TxIns {
			if input.TxId == UTxOut.TxId && input.Index == UTxOut.Index {
				exists = true
				break Outer
			}
		}
	}
	return exists

}
func makeCoinbaseTx(address string) *Tx {
	TxIns := []*TxIn{
		{"", -1, "COINBASE"},
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
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, errors.New("not enough money")
	}
	TxIns := []*TxIn{}
	TxOuts := []*TxOut{}
	totalIn := 0
	for _, UTxOut := range UTxOutsByAddress(from, Blockchain()) {
		if totalIn >= amount {
			break
		}
		TxIns = append(TxIns, &TxIn{UTxOut.TxId, UTxOut.Index, from})
		totalIn += UTxOut.Amount
	}

	change := totalIn - amount
	if change > 0 {
		changeOutput := &TxOut{from, change}
		TxOuts = append(TxOuts, changeOutput)
	}
	TxOuts = append(TxOuts, &TxOut{to, amount})

	tx := &Tx{
		Timestamp: int(time.Now().Unix()),
		TxIns:     TxIns,
		TxOuts:    TxOuts,
	}
	tx.getId()
	return tx, nil
}
func (m *mempool) AddTx(to string, amount int) error {
	Tx, err := makeTx("taeyun", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, Tx)
	return nil
}
func (m *mempool) TxsToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("taeyun")
	Txs := m.Txs
	Txs = append(Txs, coinbase)
	m.Txs = nil
	return Txs
}
