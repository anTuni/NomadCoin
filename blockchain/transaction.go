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

	var TxIns []*TxIn
	total := 0
	for _, TxOut := range Blockchain().TxOutsByAddress(from) {
		if total >= amount {
			break
		}
		TxIn := &TxIn{Owner: TxOut.Owner, Amount: TxOut.Amount}
		TxIns = append(TxIns, TxIn)
		total += TxOut.Amount
	}

	var TxOuts []*TxOut
	change := total - amount
	if change > 0 {
		TxOuts = append(TxOuts, &TxOut{Owner: from, Amount: change})
	}
	TxOuts = append(TxOuts, &TxOut{Owner: to, Amount: amount})

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
