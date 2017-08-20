package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

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
		"gender": "female", 
		"status": "single",
		"blood_type": "AB",
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

// Test get single user
func TestGetUser(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	/*var m map[string]interface{}*/
	//json.Unmarshal(response.Body.Bytes(), &m)

	/*fmt.Printf("%+v\n", m)*/
}

// Test get multiple users
func TestGetUsers(t *testing.T) {
	clearUsersTable()
	addUsers(2)

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

// Test Update Existing User
func TestUpdateUser(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{
		"name": "Test User",
		"username": "TestUser",
		"gender": "female", 
		"status": "single",
		"blood_type": "AB",
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

func TestDeleteUser(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func TestGetToken(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	a.DB.Exec("UPDATE users SET password = crypt('test123', gen_salt('md5')) where id=1")

	payload := []byte(`{
		"email": "0User@teahrm.id", 
		"password": "test123" 
	}`)

	req, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(payload))
	response := executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		t.Errorf("Status false")
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestChangeToken(t *testing.T) {
	clearUsersTable()
	addUsers(1)

	payload := []byte(`{
		"id": 1, 
		"api_token": "loremipsumdolorsitamet" 
	}`)

	req, _ := http.NewRequest("PUT", "/token", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	a.DB.Exec("UPDATE users SET password = crypt('test123', gen_salt('md5')) where id=1")

	payload = []byte(`{
		"email": "0User@teahrm.id", 
		"password": "test123" 
	}`)

	req, _ = http.NewRequest("POST", "/token", bytes.NewBuffer(payload))
	response = executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		t.Errorf("Status false")
	}

	if m["api_token"] == "loremipsumdolorsitamet" {
		t.Errorf("Token is not changed")
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec(
			"INSERT INTO users(name, username, gender, status, blood_type, email, timezone, language, signature, api_token) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",

			"User X"+strconv.Itoa(i),
			"Userx"+strconv.Itoa(i),
			"male",
			"single",
			"AB",
			strconv.Itoa(i)+"User@teahrm.id",
			"Asia/Jakarta",
			"Bahasa Indonesia",
			"Regards",
			"loremipsumdolorsitamet",
		)
	}
}

func clearUsersTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}
