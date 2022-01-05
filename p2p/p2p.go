package p2p

import (
	"net/http"

	"github.com/anTuni/NomadCoin/utils"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := Upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

}
