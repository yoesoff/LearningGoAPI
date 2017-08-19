package main

import (
	"os"

	"gitlab.com/mhyusufibrahim/teahrm/app"
	"gitlab.com/mhyusufibrahim/teahrm/public"
)

var (
	a app.App
)

func init() {
	a = app.App{}
	a.Initialize(
		os.Getenv("TEAHRM_DB_USERNAME"),
		os.Getenv("TEAHRM_DB_PASSWORD"),
		os.Getenv("TEAHRM_DB_NAME"),
	)

	a.Router.HandleFunc("/", public.HomeHandler)
	a.Router.HandleFunc("/login", public.LoginHandler)
	a.Router.HandleFunc("/register", public.RegisterHandler)

}

func main() {
	a.Run(":9090")
}
