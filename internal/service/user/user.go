package user

import (
	"context"
	"github.com/hokkung/go-groceries/internal/repository"
)

type UserService interface {
	Get(ctx context.Context, id int) (int, error)
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

func (s *User) Get(ctx context.Context, id int) (int, error) {
	return id, nil
}
