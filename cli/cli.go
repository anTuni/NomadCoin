package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/anTuni/NomadCoin/explorer"
	"github.com/anTuni/NomadCoin/rest"
)

func usage() {
	fmt.Printf("Hello this is Practice \n\n")
	fmt.Printf("Please use Commands below :\n\n")
	fmt.Printf("explorer : Start http server\n")
	fmt.Printf("rest : Start RESTful API sever\n")
	os.Exit(0)
}
func Start() {

	port := flag.Int("port", 4000, "pleas Counter")
	restport := flag.Int("restport", 3000, "pleas Counter")
	explorerport := flag.Int("explorerport", 5000, "pleas Counter")
	mode := flag.String("mode", "rest", "pleas Counter")

	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "explorer":
		explorer.Start(*port)
	case "both":
		go rest.Start(*restport)
		explorer.Start(*explorerport)
	default:
		usage()
	}
	fmt.Println(*port)
	fmt.Println(*mode)
}
