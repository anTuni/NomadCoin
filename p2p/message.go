package p2p

import (
	"encoding/json"
	"fmt"
	"strings"

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
	MessageNewBlockNotify
	MessageNewTxNotify
	MessageNewPeerNotify
)

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}
func sendAllBlocksResponse(p *peer) {
	fmt.Printf("Sending all blocks Response  to  %s\n", p.key)
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}
func sendRequestAllBlock(p *peer) {
	fmt.Printf("Sending Request of all blocks  to  %s\n", p.key)
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}
func SendNewestBlock(p *peer) {
	fmt.Printf("Sending the Newest block to  %s\n", p.key)
	block, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, block)
	p.inbox <- m
}
func sendNewBlockNotify(b *blockchain.Block, p *peer) {
	m := makeMessage(MessageNewBlockNotify, b)
	p.inbox <- m
}
func sendNewTxNotify(tx *blockchain.Tx, p *peer) {
	m := makeMessage(MessageNewTxNotify, tx)
	p.inbox <- m

}
func sendNewPeerNotify(payload string, p *peer) {
	m := makeMessage(MessageNewPeerNotify, payload)
	p.inbox <- m

}
func handelMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Recieve the Newest block from  %s\n", p.key)
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Height >= b.Height {
			sendRequestAllBlock(p)
		} else {
			SendNewestBlock(p)
		}

	case MessageAllBlocksRequest:
		fmt.Printf("Recieve the All Blocks request from  %s\n", p.key)
		sendAllBlocksResponse(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Recieve the All Blocks response from  %s\n", p.key)
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().Replace(payload)
	case MessageNewBlockNotify:
		var payload *blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().AddPeerBlock(payload)
	case MessageNewTxNotify:
		var payload *blockchain.Tx
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Mempool().AddPeerTx(payload)
	case MessageNewPeerNotify:
		var payload string
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		fmt.Printf("New peer's address is \"%s\" ", payload)
		parts := strings.Split(payload, ":")
		AddPeers(parts[0], parts[1], parts[2], false)
	}

}
