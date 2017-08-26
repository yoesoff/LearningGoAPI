package main

import (
	"os"

	"gitlab.com/mhyusufibrahim/teahrm/app"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("TEAHRM_DB_USERNAME"),
		os.Getenv("TEAHRM_DB_PASSWORD"),
		os.Getenv("TEAHRM_DB_NAME"),
	)

	a.Run(":9090")
}
