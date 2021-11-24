package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/anTuni/NomadCoin/utils"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddingBlock struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler")
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "GET",
			Description: "Get all blocks",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add a block to blockchain",
			Payload:     "data:String",
		},
	}
	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBolocks())
	case "POST":
		var addingBlock AddingBlock
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addingBlock))
		blockchain.GetBlockchain().AddBlock(addingBlock.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {

	fmt.Printf("Listen on %s", port)

	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)

	log.Fatal(http.ListenAndServe(port, nil))
}
