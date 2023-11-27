package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/hokkung/go-groceries/repository/mock"
	"github.com/hokkung/go-groceries/service"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite

	underTest service.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)

	s.underTest = service.NewUserService(mockUserRepository)
}

func (s *UserServiceTestSuite) TestGet() {
	one := 1
	res := s.underTest.Get(one)

	s.Equal(one, res)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
