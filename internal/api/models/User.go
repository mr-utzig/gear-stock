package models

import "database/sql"

type User struct {
	ID        int    `json:"id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Status    bool   `json:"status"`
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (o *UserModel) GetAllUsers() *[]User {
	return &[]User{}
}

func (o *UserModel) CreateUser() *User {
	return &User{}
}

func (o *UserModel) GetUser(id int) *User {
	return &User{}
}

func (o *UserModel) UpdateUser(id int) *User {
	return &User{}
}
