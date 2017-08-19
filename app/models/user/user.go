package user

import (
	"database/sql"
	"strings"
	"time"

	"github.com/lib/pq"
)

var (
	table         = "users"
	fields        = `id, name, username, gender, status, blood_type, email, is_active, timezone, language, signature, deleted_at, created_at, updated_at`
	fields_update = `name=$1, username=$2, gender=$3, status=$4, blood_type=$5, email=$6, is_active=$7, timezone=$8, language=$9, signature=$10, deleted_at=$11, updated_at=$12 `
	field_insert  = "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13"
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
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
}

func (u *User) GetUser(db *sql.DB) error {
	row_user := db.QueryRow(`SELECT	`+fields+` FROM `+table+` WHERE id=$1`, u.ID)

	return row_user.Scan(
		&u.ID,
		&u.Name,
		&u.Username,
		&u.Gender,
		&u.Status,
		&u.Blood_type,
		&u.Email,
		&u.Is_active,
		&u.Timezone,
		&u.Language,
		&u.Signature,
		&u.Deleted_at,
		&u.Created_at,
		&u.Updated_at)
}

func (u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec(`UPDATE `+table+` SET `+fields_update+` WHERE id=$13`,
		u.Name,
		u.Username,
		u.Gender,
		u.Status,
		u.Blood_type,
		u.Email,
		u.Is_active,
		u.Timezone,
		u.Language,
		u.Signature,
		u.Deleted_at,
		u.Updated_at,
		u.ID)

	return err
}

func (u *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM "+table+"  WHERE id=$1", u.ID)

	return err
}

func (u *User) CreateUser(db *sql.DB) error {
	fields_min_id := strings.Replace(fields, "id, ", "", -1)

	err := db.QueryRow(`INSERT INTO `+table+`(`+fields_min_id+`) VALUES(`+field_insert+`) RETURNING id`,
		u.Name,
		u.Username,
		u.Gender,
		u.Status,
		u.Blood_type,
		u.Email,
		u.Is_active,
		u.Timezone,
		u.Language,
		u.Signature,
		u.Deleted_at,
		u.Created_at,
		u.Updated_at,
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *sql.DB, start int, count int) ([]User, error) {
	rows, err := db.Query(`SELECT `+fields+` FROM `+table+` LIMIT $1 OFFSET $2`, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Username,
			&u.Gender, &u.Status, &u.Blood_type,
			&u.Email,
			&u.Is_active,
			&u.Timezone,
			&u.Language,
			&u.Signature,
			&u.Deleted_at,
			&u.Created_at,
			&u.Updated_at,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
