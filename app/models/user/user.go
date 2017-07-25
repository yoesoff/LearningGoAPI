package user

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"Email"`
	Is_active bool   `json:"is_active"`
	Timezone  string `json:"timezone"`
	Language  string `json:"language"`
	Signature string `json:"signature"`

	//https://gobyexample.com/time
	Deleted_at string `json:"deleted_at"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func (u *User) GetUser(db *sql.DB) error {
	return db.QueryRow(
		`SELECT	
			name, 
			username, 
			email, 
			is_active, 
			timezone, 
			language, 
			signature, 
			deleted_at, 
			created_at, 
			updated_at  
		FROM 
			users 
		WHERE 
			id=$1`, u.ID).Scan(&u.Name, &u.Username, &u.Email, &u.Is_active, &u.Timezone, &u.Signature, &u.Deleted_at, &u.Created_at, &u.Updated_at)

}

func (u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec(
		`UPDATE 
			users 
		SET  
			name=$1, 
			username=$2, 
			email=$3, 
			is_active=$4, 
			timezone=%5, 
			language=%6, 
			signature=%7, 
			deleted_at=%8, 
			updated_at=%9 
		WHERE 
			id=$10`,
		u.Name,
		u.Username,
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
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}

func (u *User) CreateUser(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO 
			users(name, username, email, is_active, timezone, language, signature, deleted_at, Updated_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		    RETURNING id`,
		u.Name,
		u.Username,
		u.Email,
		u.Is_active,
		u.Timezone,
		u.Language,
		u.Signature,
		u.Deleted_at,
		u.Updated_at,
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *sql.DB, start int, count int) ([]User, error) {
	rows, err := db.Query(
		`SELECT 
			name, 
			username, 
			email, 
			is_active, 
			timezone, 
			language, 
			signature, 
			deleted_at, 
			created_at, 
			updated_at 
		FROM 
			users 
		LIMIT $1 OFFSET $2`,
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			u.Name,
			u.Username,
			u.Email,
			u.Is_active,
			u.Timezone,
			u.Language,
			u.Signature,
			u.Deleted_at,
			u.Updated_at,
			u.ID,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
