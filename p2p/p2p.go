package p2p

import (
	"fmt"
	"net/http"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Conns []*websocket.Conn
var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := Upgrader.Upgrade(rw, r, nil)
	Conns = append(Conns, conn)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("Message from the CL : %s\n", p)
		message := fmt.Sprintf("Message : %s\n", p)
		for _, conn := range Conns {
			conn.WriteMessage(websocket.TextMessage, []byte(message))
		}
	}

}
