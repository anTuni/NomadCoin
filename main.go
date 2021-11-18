package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/anTuni/NomadCoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	data := homeData{"Here is Title", blockchain.GetBlockchain().AllBolocks()}
	tmpl.Execute(rw, data)
}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
