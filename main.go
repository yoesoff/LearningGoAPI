package main

import (
	"os"

	a "gitlab.com/mhyusufibrahim/teahrm/app"
)

var (
	app a.App
)

func init() {
	app = a.App{}
	app.Initialize(
		os.Getenv("TEAHRM_DB_USERNAME"),
		os.Getenv("TEAHRM_DB_PASSWORD"),
		os.Getenv("TEAHRM_DB_NAME"),
	)

	//app.Run(":9090")
}

func main() {
	/* r := mux.NewRouter()*/

	//r.HandleFunc("/", public.HomeHandler)
	//r.HandleFunc("/login", public.LoginHandler)
	//r.HandleFunc("/register", public.RegisterHandler)

	//http.ListenAndServe(":9090", r)
	app.Run(":9090")
}
