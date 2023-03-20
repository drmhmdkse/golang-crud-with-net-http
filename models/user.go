package models

type User struct {
	ID        int    `json:"id" `
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	Password  string `json:"password"`
}
