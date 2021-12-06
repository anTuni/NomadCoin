package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addingBlock struct {
	Message string
}
type errorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler")
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/block"),
			Method:      "GET",
			Description: "Get a block",
			Payload:     "id:int",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "Get all blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block to blockchain",
			Payload:     "data:String",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		return
		// json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBolocks())
	case "POST":
		return
		// var addingBlock addingBlock
		// utils.HandleErr(json.NewDecoder(r.Body).Decode(&addingBlock))
		// blockchain.GetBlockchain().AddBlock(addingBlock.Message)
		// rw.WriteHeader(http.StatusCreated)
	}
}
func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	if err == blockchain.ErrorNotFound {
		json.NewEncoder(rw).Encode(errorMessage{fmt.Sprint(err)})
	} else {
		json.NewEncoder(rw).Encode(block)
	}
}
func jsonContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
func Start(aPort int) {
	router := mux.NewRouter()

	port := fmt.Sprintf(":%d", aPort)
	fmt.Printf("Listen on %s", port)
	router.Use(jsonContentMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}
