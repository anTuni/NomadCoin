package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	key     string
	address string
	port    string
	Conn    *websocket.Conn
	inbox   chan []byte
}

func (p *peer) close() {
	delete(Peers, p.key)
}
func (p *peer) read() {
	defer p.close()
	for {
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			fmt.Println("closed from read")
			break
		}
		fmt.Printf("%s\n", m)
	}
}
func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox
		if !ok {
			fmt.Println("closed from write")
			break
		}
		p.Conn.WriteMessage(websocket.TextMessage, m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		key:     key,
		address: address,
		port:    port,
		Conn:    conn,
		inbox:   make(chan []byte),
	}
	Peers[key] = p

	go p.read()
	go p.write()
	return p
}
