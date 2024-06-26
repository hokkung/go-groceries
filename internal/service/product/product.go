package product

import (
	"github.com/hokkung/go-groceries/internal/repository"
)

type ProductService interface {
	Get(id int) int
}

type Product struct {
	productRepository repository.ProductRepository
}

func (s *Product) Get(id int) int {
	return id
}

// NewProductService creates instance
func NewProductService(productRepository repository.ProductRepository) *Product {
	return &Product{productRepository: productRepository}
}

// ProvideProductService provide instance for di
func ProvideProductService(productRepository repository.ProductRepository) *Product {
	return NewProductService(productRepository)
}
