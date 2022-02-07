package p2p

import (
	"fmt"
	"net/http"

	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	fmt.Println("Upgrade r.RemoteAddr", r.RemoteAddr)
	ip := utils.Splitter(r.RemoteAddr, ":", 0)

	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	fmt.Printf("Port : %s wants to Upgrade the conn\n", openPort)

	conn, err := Upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)

}
func AddPeers(address, port, openPort string, broadcast bool) {

	fmt.Printf("Port : %s wants to connect to Port : %s\n", openPort, port)
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	p := initPeer(conn, address, port)
	if broadcast {
		broadcastNewPeer(p)
	}
	SendNewestBlock(p)
}

func BroadcastNewBlock(b *blockchain.Block) {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	for _, p := range Peers.v {
		sendNewBlockNotify(b, p)
	}
}
func BroadcastNewTx(tx *blockchain.Tx) {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	for _, p := range Peers.v {
		sendNewTxNotify(tx, p)
	}
}
func broadcastNewPeer(newPeer *peer) {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	for key, p := range Peers.v {
		if key != newPeer.key {
			payload := fmt.Sprintf("%s:%s", newPeer.key, p.port)
			sendNewPeerNotify(payload, p)
		}
	}
}
