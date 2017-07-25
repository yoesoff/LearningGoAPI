package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	a "gitlab.com/mhyusufibrahim/teahrm/app"
	"gitlab.com/mhyusufibrahim/teahrm/public"
)

var (
	app a.App
)

func init() {
	app = a.App{}
	app.Initialize(os.Getenv("TEAHRM_DB_USERNAME"), os.Getenv("TEAHRM_DB_PASSWORD"), os.Getenv("TEAHRM_DB_NAME"))

	//app.Run(":8080")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", public.HomeHandler)
	r.HandleFunc("/login", public.LoginHandler)
	r.HandleFunc("/register", public.RegisterHandler)

	http.ListenAndServe(":9090", r)
}
