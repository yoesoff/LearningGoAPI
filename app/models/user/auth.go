package user

import (
	"database/sql"
)

type Auth struct {
	Status    bool   `json:"status"`
	ID        int    `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Api_token string `json:"api_token"`
}

// Exchange Password and email with Token
func (auth *Auth) GetApiToken(db *sql.DB) error {
	row_user := db.QueryRow(`SELECT password = crypt($1, password), id, email, api_token FROM `+table+` WHERE email=$2`, auth.Password, auth.Email)

	return row_user.Scan(
		&auth.Status,
		&auth.ID,
		&auth.Email,
		&auth.Api_token,
	)
}

// Change Token By Token and User ID
func (auth *Auth) ChangeApiToken(db *sql.DB) error {
	_, err := db.Exec(`UPDATE `+table+` SET api_token = TO_CHAR(now(), 'DD-MON-YYYY_HH24-MI-SS-US')||'_'||randomStringByLength(50) WHERE id=$1 AND api_token=$2 `, auth.ID, auth.Api_token)

	return err
}
