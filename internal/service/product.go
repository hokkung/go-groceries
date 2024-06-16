package service

import (
	"github.com/hokkung/go-groceries/internal/repository"
)

type ProductService interface {
	Get(id int) int
}

type productService struct {
	productRepository repository.ProductRepository
}

func (s *productService) Get(id int) int {
	return id
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func ProvideProductService(productRepository repository.ProductRepository) ProductService {
	return NewProductService(productRepository)
}
