package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	app "gitlab.com/mhyusufibrahim/teahrm/app"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize(os.Getenv("TEAHRM_DB_USERNAME"), os.Getenv("TEAHRM_DB_PASSWORD"), os.Getenv("TEAHRM_DB_TEST_NAME"))

	pretest()

	code := m.Run()

	posttest()

	os.Exit(code)
}

func pretest() {
	var args = []string{
		os.Getenv("TEAHRM_DB_SERVER"),
		"\"user=" + os.Getenv("TEAHRM_DB_USERNAME") + " dbname=" + os.Getenv("TEAHRM_DB_TEST_NAME") + " sslmode=disable\"",
		"up",
	}

	exe := exec.Command("goose", args...)
	exe.Dir = os.Getenv("TEAHRM_DB_MIGRATION")
	/* result, err := exe.Output()*/
	//fmt.Println(string(result))
	/*fmt.Println(err)*/

	output := exe.Run()
	fmt.Println(output)
}

func posttest() {
}
