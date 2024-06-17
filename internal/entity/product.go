package entity

import (
	"gorm.io/gorm"

	"github.com/hokkung/go-groceries/pkg/domain"
)

type Product struct {
	gorm.Model
	domain.Name
}

func (m *Product) Table() string {
	return "products"
}
