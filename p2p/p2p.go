package p2p

import (
	"fmt"
	"net/http"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := Upgrader.Upgrade(rw, r, nil)
	initPeer(conn, "XX", "XX")

	utils.HandleErr(err)
}
func AddPeers(address, port string) {

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws", address, port), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
