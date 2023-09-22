package repository

import "gorm.io/gorm"

type ProductRepository interface {
	Model() string
}

type productRepository struct {
	DB *gorm.DB
}

func (r *productRepository) Model() string {
	return "products"
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

func ProvideProductRepository(db *gorm.DB) (ProductRepository, func(), error) {
	return NewProductRepository(db), func(){}, nil
}
