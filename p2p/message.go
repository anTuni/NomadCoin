package p2p

import (
	"encoding/json"
	"fmt"

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

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}
func SendNewestBlock(p *peer) {
	block, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, block)
	p.inbox <- m
}

func handelMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var b blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &b))
		fmt.Printf("From Peer:%s Message payload %s Kind of : %d", p.key, b.Hash, m.Kind)
	}

}
