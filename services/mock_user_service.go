package services

import (
	"todo-case/models"

)
type UserService interface {
	ValidateCredentials(username, password string) (*models.User, bool)
}

type MockUserService struct {
	users []models.User
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		users: []models.User{
			{ID: 1, Username: "admin", Password: "1234",UserType: "admin",},
			{ID: 2, Username: "user1", Password: "test",UserType: "user",},
			{ID: 3, Username: "user2", Password: "user",UserType: "user",},
			{ID: 4, Username: "user3", Password: "guest",UserType: "user",},
		},
	}
}

func (s *MockUserService) ValidateCredentials(username, password string) (*models.User, bool) {
	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return &user, true
		}
	}
	return nil, false
}
