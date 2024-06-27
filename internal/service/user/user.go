package user

import (
	"github.com/hokkung/go-groceries/internal/repository"
)

type UserService interface {
	Get(id int) int
}

type User struct {
	userRepository repository.UserRepository
}

// NewUserService creates instance
func NewUserService(userRepository repository.UserRepository) *User {
	return &User{userRepository: userRepository}
}

// ProvideUserService provides instance for di
func ProvideUserService(userRepository repository.UserRepository) *User {
	return NewUserService(userRepository)
}

func (s *User) Get(id int) int {
	return id
}
