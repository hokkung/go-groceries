package repository

import "gorm.io/gorm"

//go:generate mockgen -source ./product.go -destination ./mock/mock_product.go
type ProductRepository interface {
}

type Product struct {
	DB *gorm.DB
}

// NewProductRepository creates instance
func NewProductRepository(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

// ProvideProductRepository provides instance for di
func ProvideProductRepository(db *gorm.DB) (ProductRepository, func(), error) {
	return NewProductRepository(db), func() {}, nil
}
