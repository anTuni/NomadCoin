package blockchain

import (
	"errors"
	"sync"
	"time"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/anTuni/NomadCoin/wallet"
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

type TxIn struct {
	TxId      string `json:"txid"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}

type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}
type UTxOut struct {
	TxId   string `json:"txid"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

type mempool struct {
	Txs map[string]*Tx
	m   sync.Mutex
}

var m *mempool
var memOnce sync.Once

func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{
			Txs: make(map[string]*Tx),
		}
	})
	return m
}

func isOnMempool(UTxOut *UTxOut) bool {
	exists := false
Outer:
	for _, Tx := range Mempool().Txs {
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

var ErrorNoMoney = errors.New("not enough money")
var ErrorNotValid = errors.New("not valid Transaction")

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, ErrorNoMoney
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
	tx.sign()
	valid := validate(tx)
	if !valid {
		return nil, ErrorNotValid
	}
	return tx, nil
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func (t *Tx) sign() {
	for _, TxIn := range t.TxIns {
		TxIn.Signature = wallet.Sign(t.Id, wallet.Wallet())
	}
}

func validate(tx *Tx) bool {
	valid := true
	for _, TxIn := range tx.TxIns {
		prevTx := FindTx(Blockchain(), TxIn.TxId)
		if prevTx == nil {
			valid = false
			break
		}
		valid = wallet.Verify(TxIn.Signature, tx.Id, prevTx.TxOuts[TxIn.Index].Address)
		if !valid {
			return valid
		}
	}
	return valid
}
func (m *mempool) AddTx(to string, amount int) (*Tx, error) {
	Tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return nil, err
	}
	m.Txs[Tx.Id] = Tx
	return Tx, nil
}
func (m *mempool) TxsToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)
	var Txs []*Tx
	for _, tx := range m.Txs {
		Txs = append(Txs, tx)
	}
	Txs = append(Txs, coinbase)
	m.Txs = make(map[string]*Tx)
	return Txs
}

func (m *mempool) AddPeerTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()
	m.Txs[tx.Id] = tx

}
