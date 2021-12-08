package main

import (
	"github.com/anTuni/NomadCoin/cli"
	"github.com/anTuni/NomadCoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
