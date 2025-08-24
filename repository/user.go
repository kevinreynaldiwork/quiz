package repository

import (
	"Quiz/structs"
	"database/sql"
	"errors"
)

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var user structs.User
	err := db.QueryRow(
		`SELECT id, username, password, created_at, created_by, modified_at, modified_by 
		 FROM users WHERE username=$1`,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.ModifiedAt,
		&user.ModifiedBy,
	)

	if err == sql.ErrNoRows {
		return user, errors.New("user not found")
	}
	return user, err
}

func InsertUser(db *sql.DB, user structs.User) error {
	_, err := db.Exec(`
		INSERT INTO users (username, password, created_at, created_by, modified_at, modified_by) 
		VALUES ($1, $2, NOW(), $3, NOW(), $3)`,
		user.Username, user.Password, user.CreatedBy,
	)
	return err
}
