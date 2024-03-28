package models

type User struct {
	ID      string
	Name    string `json:"name"`
	Age     int
	Address string
}
