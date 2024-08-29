package repository

import (
	"context"
	"github.com/hokkung/go-groceries/internal/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source ./user.go -destination ./mock/mock_user.go
type UserRepository interface {
	Model() string
	FindByName(ctx context.Context, name string) (entity.User, error)
}

type userRepository struct {
	*Base
}

func (r *userRepository) Model() string {
	return "users"
}

func (r *userRepository) FindByName(ctx context.Context, name string) (entity.User, error) {
	var user entity.User
	err := r.getDB(ctx).Model(&entity.User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		Base: NewBase(db),
	}
}

func ProvideUserRepository(db *gorm.DB) (UserRepository, func(), error) {
	return NewUserRepository(db), func() {}, nil
}
