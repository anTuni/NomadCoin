package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/anTuni/NomadCoin/blockchain"
)

const (
	port        string = ":4000"
	templateDIr string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}

}
func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBolocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDIr + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDIr + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
