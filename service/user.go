package service

import "github.com/hokkung/go-groceries/repository"

type UserService interface {
	Get(id int) int
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Get(id int) int {
	return id
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func ProvideUserService(userRepository repository.UserRepository) UserService {
	return NewUserService(userRepository)
}
