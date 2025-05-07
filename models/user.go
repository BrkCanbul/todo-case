package models

type User struct {
	Username string `json:"user_name" example:"user name"`
	Password string `json:"password" example:"password123"`
	UserType string `json:"user_type" example:"admin"`
}
