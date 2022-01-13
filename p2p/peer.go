package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	Conn *websocket.Conn
}

func (p *peer) read() {
	for {
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) {
	p := &peer{Conn: conn}
	go p.read()
	key := fmt.Sprintf("%s:%s", address, port)
	Peers[key] = p
}
