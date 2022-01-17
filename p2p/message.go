package p2p

import (
	"encoding/json"

	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/anTuni/NomadCoin/utils"
)

type MessageKind int

type Message struct {
	Kind    MessageKind
	Payload []byte
}

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

func (m *Message) addPayload(p interface{}) {
	payload, err := json.Marshal(p)
	utils.HandleErr(err)
	m.Payload = payload
}
func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}
	m.addPayload(payload)
	mJSON, err := json.Marshal(m)
	utils.HandleErr(err)
	return mJSON
}
func SendNewestBlock(p *peer) {
	block, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, block)
	p.inbox <- m
}
