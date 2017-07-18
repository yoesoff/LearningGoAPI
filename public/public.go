package public

import (
	"net/http"

	"github.com/flosch/pongo2"
)

var (
	templateFolder = "templates/"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	template := templateFolder + "public/home.html"
	home_html := pongo2.Must(pongo2.FromFile(template))

	err := home_html.ExecuteWriter(pongo2.Context{"title": "Hello World mate.", "greating": "Hai, Hello world!"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	template := templateFolder + "public/login.html"
	home_html := pongo2.Must(pongo2.FromFile(template))

	err := home_html.ExecuteWriter(pongo2.Context{"title": "Login.", "greating": "Hai, Hello world Login!"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	template := templateFolder + "public/register.html"
	home_html := pongo2.Must(pongo2.FromFile(template))

	err := home_html.ExecuteWriter(pongo2.Context{"title": "Register.", "greating": "Hai, Hello world register!"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
