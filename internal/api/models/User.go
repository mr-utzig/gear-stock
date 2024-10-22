package models

import (
	"github.com/mr-utzig/gear-stock/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `json:"id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    bool   `json:"status"`
}

type PostUser struct {
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (o *UserModel) GetAllUsers() *[]User {
	return &[]User{}
}

func (o *UserModel) CreateUser(data *PostUser) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	stmt, err := database.Turso.Prepare("INSERT INTO user (name, email, password, profile_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.Name, data.Email, string(password), data.ProfileID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        id,
		ProfileID: data.ProfileID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Status:    true,
	}

	return user, nil
}

func (o *UserModel) GetUser(id int) *User {
	return &User{}
}

func (o *UserModel) UpdateUser(id int) *User {
	return &User{}
}
