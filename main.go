package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Hello this is Practice \n\n")
	fmt.Printf("Please use Commands below :\n\n")
	fmt.Printf("explorer : Start http server\n")
	fmt.Printf("rest : Start RESTful API sever\n")
	os.Exit(0)
}
func main() {
	if len(os.Args) < 2 {
		usage()
	}
	switch os.Args[1] {
	case "explorer":
		fmt.Printf("Start explorer server\n")
	case "rest":
		fmt.Printf("Start RESTful API server\n")
	default:
		usage()
	}
}
