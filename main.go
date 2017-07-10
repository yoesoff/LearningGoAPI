package main

import (
	"net/http"

	pongo "github.com/flosch/pongo2"
	"github.com/gorilla/mux"
)

var (
	templateFolder = "templates/"
	assetsFolder   = "assets/"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":80", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template := templateFolder + "public/home.html"
	home_html := pongo.Must(pongo.FromFile(template))

	err := home_html.ExecuteWriter(pongo.Context{"title": "Hello World mate.", "greating": "Hai, Hello world!"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
