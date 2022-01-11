package p2p

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := Upgrader.Upgrade(rw, r, nil)
	address := strings.Split(r.RemoteAddr, ":")
	openPort := r.URL.Query().Get("openPort")
	initPeer(conn, address[0], openPort)

	utils.HandleErr(err)
}
func AddPeers(address, port, openPort string) {

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
