package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anTuni/NomadCoin/blockchain"
	"github.com/anTuni/NomadCoin/p2p"
	"github.com/anTuni/NomadCoin/utils"
	"github.com/anTuni/NomadCoin/wallet"
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

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}
type addTxPayloads struct {
	To     string
	Amount int
}
type addPeerPayload struct {
	Address string
	Port    string
}
type AddressResponse struct {
	Address string `json:"address"`
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
			URL:         url("/status"),
			Method:      "GET",
			Description: "Get blockchain status",
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
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get Txs by owner",
		},
		{
			URL:         url("/ws"),
			Method:      "GET",
			Description: "Upgrade HTTP to WS",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain()))
	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
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
func status(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain()))
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.BalanceByAddress(address, blockchain.Blockchain())
		utils.HandleErr(json.NewEncoder(rw).Encode(balanceResponse{address, amount}))
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(address, blockchain.Blockchain())))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
}
func transaction(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayloads
	json.NewDecoder(r.Body).Decode(&payload)
	fmt.Println("payload start")
	fmt.Print(payload)
	fmt.Println("payload end")
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
func myWallet(rw http.ResponseWriter, r *http.Request) {
	address := AddressResponse{Address: wallet.Wallet().Address}
	json.NewEncoder(rw).Encode(address)
}
func peers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var payload addPeerPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		utils.HandleErr(err)
		p2p.AddPeers(payload.Address, payload.Port, port)
		rw.WriteHeader(http.StatusOK)
	case "GET":
		json.NewEncoder(rw).Encode(p2p.AllPeers(&p2p.Peers))
	}
}
func jsonContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		next.ServeHTTP(rw, r)
	})
}
func Start(aPort int) {
	router := mux.NewRouter()

	port = fmt.Sprintf(":%d", aPort)
	fmt.Printf("Listen on %s\n", port)
	router.Use(jsonContentMiddleware, loggerMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/block/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	router.HandleFunc("/wallet", myWallet).Methods("GET")
	router.HandleFunc("/mempool", mempool).Methods("GET")
	router.HandleFunc("/ws", p2p.Upgrade).Methods("GET")
	router.HandleFunc("/peers", peers).Methods("GET", "POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}
