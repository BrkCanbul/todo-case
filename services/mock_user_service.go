package services

type User struct {
	ID       uint
	Username string
	Password string 
}

type UserService interface {
	ValidateCredentials(username, password string) (*User, bool)
}

type MockUserService struct {
	users []User
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		users: []User{
			{ID: 1, Username: "admin", Password: "1234"},
			{ID: 2, Username: "test", Password: "test"},
		},
	}
}

func (s *MockUserService) ValidateCredentials(username, password string) (*User, bool) {
	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return &user, true
		}
	}
	return nil, false
}
