package handler

import (
	"github.com/hokkung/go-groceries/internal/service"
)

type ProductHandler interface {
	Get(id int) int
}

type productHandler struct {
	productService service.ProductService
}

func (h *productHandler) Get(id int) int {
	return h.productService.Get(id)
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func ProvideProductHandler(productService service.ProductService) ProductHandler {
	return NewProductHandler(productService)
}
