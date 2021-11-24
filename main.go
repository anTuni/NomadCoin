package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
			Method:      "Post",
			Description: "add a block to blockchain",
			Payload:     "data:String",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	fmt.Println(data)

	json.NewEncoder(rw).Encode(data)
}

func main() {

	fmt.Printf("Listen on %s", port)

	http.HandleFunc("/", documentation)

	log.Fatal(http.ListenAndServe(port, nil))
}
