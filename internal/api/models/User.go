package models

type User struct {
	ID        int    `json:"id"`
	ProfileID int    `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Status    bool   `json:"status"`
}
