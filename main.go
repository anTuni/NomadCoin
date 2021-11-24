package main

import (
	"github.com/anTuni/NomadCoin/explorer"
	"github.com/anTuni/NomadCoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
