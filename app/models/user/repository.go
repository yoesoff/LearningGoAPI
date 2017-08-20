package user

import (
	"database/sql"
	"strings"
)

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
	fields2 := strings.Replace(fields, "id, ", "", -1)
	fields2 = strings.Replace(fields2, ", deleted_at, created_at, updated_at", "", -1)
	err := db.QueryRow(`INSERT INTO `+table+`(`+fields2+`) VALUES(`+fields_insert+`) RETURNING id`,
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
