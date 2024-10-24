package models

import (
	"fmt"

	"github.com/mr-utzig/gear-stock/internal/api/utils"
	"github.com/mr-utzig/gear-stock/internal/database"
	"github.com/sethvargo/go-password/password"
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

type CreateUserRequest struct {
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (o *UserModel) GetAllUsers() *[]User {
	return &[]User{}
}

func (o *UserModel) CreateUser(data *CreateUserRequest) (*User, error) {
	randpassword, err := password.Generate(12, 4, 4, false, true)
	if err != nil {
		return nil, err
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(randpassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	stmt, err := database.Turso.Prepare("INSERT INTO user (name, email, password, profile_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.Name, data.Email, string(hashpassword), data.ProfileID)
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
		Password:  string(hashpassword),
		Status:    true,
	}

	utils.SendMail(data.Email, "Credentials Created", fmt.Sprintf("Email: %v\nPassword: %v", data.Email, randpassword))

	return user, nil
}

func (o *UserModel) GetUser(id int) *User {
	return &User{}
}

func (o *UserModel) UpdateUser(id int) *User {
	return &User{}
}
