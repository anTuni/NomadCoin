package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	Conn  *websocket.Conn
	Inbox chan []byte
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
func (p *peer) write() {
	for {
		m := <-p.Inbox
		p.Conn.WriteMessage(websocket.TextMessage, m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	p := &peer{Conn: conn, Inbox: make(chan []byte)}
	go p.read()
	go p.write()
	key := fmt.Sprintf("%s:%s", address, port)
	Peers[key] = p
	return p
}
