package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/paulbellamy/ratecounter"
	app "gitlab.com/mhyusufibrahim/teahrm/app"
)

var (
	a                       = app.App{}
	counter                 = ratecounter.NewRateCounter(1 * time.Second)
	goose_connection_string = "user=" + os.Getenv("TEAHRM_DB_USERNAME") + " password=" + os.Getenv("TEAHRM_DB_PASSWORD") + " " + " dbname=" + os.Getenv("TEAHRM_DB_TEST_NAME") + " sslmode=disable"
	goose_args              = []string{
		os.Getenv("TEAHRM_DB_SERVER"),
		goose_connection_string,
	}
)

func TestMain(m *testing.M) {
	a.Initialize(os.Getenv("TEAHRM_DB_USERNAME"), os.Getenv("TEAHRM_DB_PASSWORD"), os.Getenv("TEAHRM_DB_TEST_NAME"))

	pretest()
	code := m.Run()

	defer os.Exit(code)
	defer posttest()
}

func pretest() {
	counter.Incr(1) // Start stopwatch

	fmt.Printf(`
______________________ ____________________.___ _______    ________ 
\__    ___/\_   _____//   _____/\__    ___/|   |\      \  /  _____/ 
  |    |    |    __)_ \_____  \   |    |   |   |/   |   \/   \  ___ 
  |    |    |        \/        \  |    |   |   /    |    \    \_\  \
  |____|   /_______  /_______  /  |____|   |___\____|__  /\______  /
                   \/        \/                        \/        \/ 	
	`)

	fmt.Printf(`
Testing Preparation 
	- Database: ` + os.Getenv("TEAHRM_DB_SERVER") + `
	- Database Username: ` + os.Getenv("TEAHRM_DB_USERNAME") + `
	- DB Test: ` + os.Getenv("TEAHRM_DB_TEST_NAME") + `
	- Migration path: ` + os.Getenv("TEAHRM_DB_MIGRATION") + `\n`)

	fmt.Println("\n\nRunning Migration Up\n")

	goose_args = append(goose_args, "up")
	exe := exec.Command("goose", goose_args...)
	exe.Dir = os.Getenv("TEAHRM_DB_MIGRATION")
	result, err := exe.Output()
	fmt.Println(string(result))
	fmt.Println(err)

	/*output := exe.Run()*/
	/*fmt.Println(output)*/
}

func posttest() {
	fmt.Println("\nTesting is Done, Currently Running Migration Reset\n")

	goose_args[2] = "reset"
	exe := exec.Command("goose", goose_args...)
	exe.Dir = os.Getenv("TEAHRM_DB_MIGRATION")
	result, err := exe.Output()
	fmt.Println(string(result))
	fmt.Println(err)

	fmt.Println(counter.Rate()) // End stopwatch
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
