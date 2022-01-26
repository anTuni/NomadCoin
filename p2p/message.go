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
	MessageNewBlockNotify
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
		fmt.Print("MessageNewBlockNotify")
	}

}
