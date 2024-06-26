package product

import (
	"github.com/hokkung/go-groceries/internal/service/product"
)

type Product struct {
	productService product.ProductService
}

// NewProductHandler creates instance
func NewProductHandler(productService product.ProductService) *Product {
	return &Product{productService: productService}
}

// ProvideProductHandler provides instance for di
func ProvideProductHandler(productService product.ProductService) *Product {
	return NewProductHandler(productService)
}

func (h *Product) Get(id int) int {
	return h.productService.Get(id)
}
