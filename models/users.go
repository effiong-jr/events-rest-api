package models

import (
	"example.com/events-rest-api/db"

	"example.com/events-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" validate:"min=8"`
	// ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

func (u *User) RegisterUser() (*User, error) {

	query := `INSERT INTO users (email, password)
	VALUES (?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(&u.Email, &hashedPassword)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	u.ID = id
	u.Password = hashedPassword

	return u, err
}

func (u *User) LoginUser() (*User, error) {

	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string

	if err := row.Scan(&u.ID, &hashedPassword); err != nil {
		return nil, err
	}

	err := utils.CompareHashedPassword(hashedPassword, u.Password)

	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword

	return u, err

}
