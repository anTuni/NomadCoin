package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/anTuni/NomadCoin/blockchain"
)

const (
	templateDIr string = "explorer/templates/"
)

var templates *template.Template

// type homeData struct {
// 	PageTitle string
// 	Blocks    []*blockchain.Block
// }

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}

}
func home(rw http.ResponseWriter, r *http.Request) {
	return
	// data := homeData{"Home", blockchain.GetBlockchain().AllBolocks()}
	// templates.ExecuteTemplate(rw, "home", data)
}

func Start(port int) {
	templates = template.Must(template.ParseGlob(templateDIr + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDIr + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
