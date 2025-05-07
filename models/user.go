package models

type User struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"user_name" example:"user name"`
	Password string `json:"password" example:"password123"`
	UserType string `json:"user_type" example:"admin"`
}
