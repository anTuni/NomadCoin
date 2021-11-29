package main

import (
	"flag"
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

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Set the port of server")

	switch os.Args[1] {
	case "explorer":
		fmt.Printf("Start explorer server\n")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(*portFlag)
	}
}
