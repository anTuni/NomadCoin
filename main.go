package main

import (
	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/anTuni/NomadCoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
