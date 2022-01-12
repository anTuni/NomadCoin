package p2p

import (
	"fmt"
	"net/http"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)

	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}

	conn, err := Upgrader.Upgrade(rw, r, nil)
	initPeer(conn, ip, openPort)
	utils.HandleErr(err)
}
func AddPeers(address, port, openPort string) {

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?osspenPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
