package user

import (
	"github.com/lib/pq"
)

var (
	table         = "users"
	fields        = `id, name, username, gender, status, blood_type, email, is_active, timezone, language, signature, deleted_at, created_at, updated_at`
	fields_update = `name=$1, username=$2, gender=$3, status=$4, blood_type=$5, email=$6, is_active=$7, timezone=$8, language=$9, signature=$10, deleted_at=$11, updated_at=$12 `
	fields_insert = "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Gender     string `json:"gender"`
	Status     string `json:"status"`
	Blood_type string `json:"blood_type"`
	Email      string `json:"Email"`
	Is_active  bool   `json:"is_active"`
	Timezone   string `json:"timezone"`
	Language   string `json:"language"`
	Signature  string `json:"signature"`

	//https://gobyexample.com/time
	Deleted_at pq.NullTime `json:"deleted_at"`
	Created_at pq.NullTime `json:"created_at"`
	Updated_at pq.NullTime `json:"updated_at"`
}
