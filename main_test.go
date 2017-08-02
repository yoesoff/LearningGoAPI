package main_test

import (
	"encoding/json"
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

func TestGetEmptyUserTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

/*func TestCreateProduct(t *testing.T) {*/
//clearTable()

//payload := []byte(`{"name":"test product","price":11.22}`)

//req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
//response := executeRequest(req)

//checkResponseCode(t, http.StatusCreated, response.Code)

//var m map[string]interface{}
//json.Unmarshal(response.Body.Bytes(), &m)

//if m["name"] != "test product" {
//t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
//}

//if m["price"] != 11.22 {
//t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
//}

//// the id is compared to 1.0 because JSON unmarshaling converts numbers to
//// floats, when the target is a map[string]interface{}
//if m["id"] != 1.0 {
//t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
//}
/*}*/

func pretest() {
	counter.Incr(1) // Start stopwatch

	fmt.Printf(`Testing Preparation 
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
	fmt.Println("\nRunning Migration Reset\n")

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
