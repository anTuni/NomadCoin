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
	fmt.Printf("Port : %s wants to Upgrade the conn\n", openPort)

	conn, err := Upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)

}
func AddPeers(address, port, openPort string) {

	fmt.Printf("Port : %s wants to connect to Port : %s\n", openPort, port)
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	p := initPeer(conn, address, port)
	SendNewestBlock(p)
}
