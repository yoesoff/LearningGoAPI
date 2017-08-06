package main_test

import (
	"bytes"
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

// Test Empty Users
func TestGetEmptyUserTable(t *testing.T) {
	clearUsersTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

// Test Non Existed User
func TestGetNonExistentUser(t *testing.T) {
	clearUsersTable()

	req, _ := http.NewRequest("GET", "/users/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

// Test Create New User
func TestCreateUser(t *testing.T) {
	clearUsersTable()

	payload := []byte(`{
		"name": "Test User", 
		"username": "TestUser", 
		"email": "testuser@teahrm.id", 
		"is_active": true, 
		"timezone": "Asia/Jakarta", 
		"language": "Bahasa Indonesia", 
		"signature": "Regards" 
	}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "Test User" {
		t.Errorf("Expected user name to be 'Test User'. Got '%v'", m["name"])
	}
}

func TestGetUser(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	/* var m map[string]interface{}*/
	/*json.Unmarshal(response.Body.Bytes(), &m)*/

	//fmt.Printf("%+v\n", m)
}

// Test Update Existing User
func TestUpdateUser(t *testing.T) {
	clearUsersTable()
	addUsers(22)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{
		"name": "Test User",
		"username": "TestUser",
		"email": "testuser@teahrm.id",
		"is_active": true,
		"timezone": "Asia/Jakarta",
		"language": "Bahasa Indonesia",
		"signature": "Regards"
	}`)

	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
	}
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec(
			"INSERT INTO users(name, username, email, timezone, language, signature) VALUES($1, $2, $3, $4, $5, $6)",

			"User X",
			"User",
			"User@teahrm.id",
			"Asia/Jakarta",
			"Bahasa Indonesia",
			"Regards",
		)

	}
}

func clearUsersTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
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
