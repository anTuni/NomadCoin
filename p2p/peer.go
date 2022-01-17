package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

var Peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

func AllPeers(p *peers) []string {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	var peers []string
	for key := range p.v {
		peers = append(peers, key)
	}
	return peers
}
func (p *peer) close() {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	p.conn.Close()
	delete(Peers.v, p.key)
}
func (p *peer) read() {
	defer p.close()
	for {
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("closed from read")
			break
		}
		fmt.Print(m.Kind)
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
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		key:     key,
		address: address,
		port:    port,
		conn:    conn,
		inbox:   make(chan []byte),
	}
	Peers.v[key] = p

	go p.read()
	go p.write()
	return p
}
