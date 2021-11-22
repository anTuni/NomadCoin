package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
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
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func main() {

	fmt.Printf("Listen on %s", port)

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(port, nil))
}
