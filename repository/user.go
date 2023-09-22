package repository

import "gorm.io/gorm"

type UserRepository interface {
	Model() string
}

type userRepository struct {
	DB *gorm.DB
}

func (r *userRepository) Model() string {
	return "users"
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func ProvideUserRepository(db *gorm.DB) (UserRepository, func(), error) {
	return NewUserRepository(db), func(){}, nil
}
