package models

import (
	"fmt"

	"github.com/mr-utzig/gear-stock/internal/api/utils"
	"github.com/mr-utzig/gear-stock/internal/database"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64  `json:"id"`
	ProfileID    int    `json:"profile_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	SessionToken *string
	Status       bool `json:"status"`
	FirstAccess  bool
}

type UserDataResponse struct {
	ID        int64  `json:"id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
}

type UserCreateRequest struct {
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return new(UserModel)
}

func (u *UserModel) GetAllUsers() ([]UserDataResponse, error) {
	rows, err := database.Turso.Query("SELECT user_id, name, email, status, profile_id FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []UserDataResponse{}
	for rows.Next() {
		user := UserDataResponse{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Status, &user.ProfileID)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserModel) CreateUser(data *UserCreateRequest) (*User, error) {
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

	exec, err := stmt.Exec(data.Name, data.Email, string(hashpassword), data.ProfileID)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
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

func (u *UserModel) ValidateCredentials(data *UserLoginRequest) (*User, error) {
	user := new(User)

	err := database.Turso.QueryRow(
		`SELECT user_id, name, email, password, status, profile_id, session_token, first_access
		FROM user
		WHERE email=?`,
		data.Email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.ProfileID,
		&user.SessionToken,
		&user.FirstAccess,
	)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserModel) GetUser(id int) (*User, error) {
	user := new(User)
	err := database.Turso.QueryRow(
		`SELECT user_id, name, email, password, status, profile_id, session_token, first_access
		FROM user
		WHERE user_id=?`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.ProfileID,
		&user.SessionToken,
		&user.FirstAccess,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserModel) UpdateUser(id int, data *UserUpdateRequest) (*User, error) {
	stmt, err := database.Turso.Prepare("UPDATE user SET name=?, email=?, status=?, profile_id=? WHERE user_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Email, data.Status, data.ProfileID, id)
	if err != nil {
		return nil, err
	}

	user, err := u.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserModel) UpdatePassword(id int, password string) error {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	stmt, err := database.Turso.Prepare("UPDATE user SET password=? WHERE user_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(string(hashpassword), id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) UpdateSessionToken(id int64, token *string) error {
	stmt, err := database.Turso.Prepare("UPDATE user SET session_token=?, first_access=0 WHERE user_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token, id)
	if err != nil {
		return err
	}

	return nil
}
