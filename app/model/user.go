package model

import (
	"database/sql"
	"fmt"
	"os"
)

type User struct {
	ID       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Salt     string
}

func (u *User) Create(db *sql.DB) error {
	rows, err := db.Query(
		`INSERT INTO user
		(name, email, password, salt)
		VALUES (?, ?, ?, ?)`,
		u.Name, u.Email, u.Password, u.Salt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	return nil
}

func (u *User) Modify(db *sql.DB) error {
	return nil
}

func (u *User) Validate(db *sql.DB) error {
	err := db.QueryRow(
		`SELECT id, name
		FROM user
		WHERE email=?
			AND password=?`,
		u.Email,
		u.Password,
	).Scan(&u.ID, &u.Name)

	if err != nil {
		return err
	}

	return nil
}
