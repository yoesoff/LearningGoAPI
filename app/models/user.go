package model

import (
	"database/sql"
)

type user struct {
	Name      string
	Username  string
	Email     string
	Is_active bool
	Timezone  string
	Language  string
	Signature string

	//https://gobyexample.com/time
	Deleted_at string
	Created_at string
	Updated_at string
}

func (u *user) getUser(db *sql.DB) error {
	db.QueryRow(
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
			id=$1`, u.ID)

	db.Scan(&u.Name, &u.Username, &u.Email, &u.Is_active, &u.Timezone, &u.Signature, &u.Deleted_at, &u.Created_at, &u.Updated_at)

	return db
}

func (u *user) updateUser(db *sql.DB) error {
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
		p.Name,
		p.Username,
		p.Email,
		p.Is_active,
		p.Timezone,
		p.Language,
		p.Signature,
		p.Deleted_at,
		p.Updated_at,
		p.ID)

	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", p.ID)

	return err
}

func (u *user) createUser(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO 
			users(name, price) 
		VALUES($1, $2) 
		    RETURNING id`,
		p.Name,
		p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
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

	users := []user{}

	for rows.Next() {
		var p user
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}
