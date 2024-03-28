package models

type Auth struct {
	ID       string
	Email    string `json:"email"`
	Password string `json:"password"`
}
