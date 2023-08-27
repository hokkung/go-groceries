package entity

import (
	"github.com/hokkung/go-groceries/pkg/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	domain.Name
}
