package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/paulbellamy/ratecounter"
	app "gitlab.com/mhyusufibrahim/teahrm/app"
)

var (
	a                       app.App
	counter                 = ratecounter.NewRateCounter(1 * time.Second)
	goose_connection_string = "user=" + os.Getenv("TEAHRM_DB_USERNAME") + " password=" + os.Getenv("TEAHRM_DB_PASSWORD") + " " + " dbname=" + os.Getenv("TEAHRM_DB_TEST_NAME") + " sslmode=disable"
	goose_args              = []string{
		os.Getenv("TEAHRM_DB_SERVER"),
		goose_connection_string,
	}
)

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize(os.Getenv("TEAHRM_DB_USERNAME"), os.Getenv("TEAHRM_DB_PASSWORD"), os.Getenv("TEAHRM_DB_TEST_NAME"))

	pretest()

	code := m.Run()

	posttest()

	os.Exit(code)
}

func pretest() {
	counter.Incr(1) // Start stopwatch

	fmt.Println("Testing Preparation")
	fmt.Println("- Database: " + os.Getenv("TEAHRM_DB_SERVER"))
	fmt.Println("- Database Username: " + os.Getenv("TEAHRM_DB_USERNAME"))
	fmt.Println("- DB Test: " + os.Getenv("TEAHRM_DB_TEST_NAME"))
	fmt.Println("- Migration path: " + os.Getenv("TEAHRM_DB_MIGRATION"))

	fmt.Println("\nRunning Migration Up\n")

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
	fmt.Println("\nRunning Migration Reset\n")

	goose_args[2] = "reset"
	exe := exec.Command("goose", goose_args...)
	exe.Dir = os.Getenv("TEAHRM_DB_MIGRATION")
	result, err := exe.Output()
	fmt.Println(string(result))
	fmt.Println(err)

	fmt.Println(counter.Rate()) // End stopwatch
}
