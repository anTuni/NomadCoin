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
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("Message from the CL : %s\n", p)
		message := fmt.Sprintf("Message : %s\n", p)
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}

}
