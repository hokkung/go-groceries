package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/hokkung/go-groceries/repository/mock"
	"github.com/hokkung/go-groceries/service"
	"github.com/stretchr/testify/suite"
)

type ProductServiceTestSuite struct {
	suite.Suite

	underTest service.ProductService
}

func (s *ProductServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	mockProductRepository := mock_repository.NewMockProductRepository(ctrl)

	s.underTest = service.NewProductService(mockProductRepository)
}

func (s *ProductServiceTestSuite) TestGet() {
	one := 1
	res := s.underTest.Get(one)

	s.Equal(one, res)
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
