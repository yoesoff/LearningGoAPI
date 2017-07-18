package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/mhyusufibrahim/teahrm/public"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", public.HomeHandler)
	r.HandleFunc("/login", public.LoginHandler)
	r.HandleFunc("/register", public.RegisterHandler)

	http.ListenAndServe(":9090", r)
}
