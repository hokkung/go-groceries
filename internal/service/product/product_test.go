package product_test

import (
	"github.com/hokkung/go-groceries/internal/repository/mock"
	"github.com/hokkung/go-groceries/internal/service/product"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ProductServiceTestSuite struct {
	suite.Suite

	underTest product.ProductService
}

func (s *ProductServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	mockProductRepository := mock_repository.NewMockProductRepository(ctrl)

	s.underTest = product.ProvideProductService(mockProductRepository)
}

func (s *ProductServiceTestSuite) TestGet() {
	one := 1
	res := s.underTest.Get(one)

	s.Equal(one, res)
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
