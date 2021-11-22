package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anTuni/NomadCoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler")
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	byte, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", byte)
}

func main() {

	fmt.Printf("Listen on %s", port)

	http.HandleFunc("/add", home)

	log.Fatal(http.ListenAndServe(port, nil))
}
