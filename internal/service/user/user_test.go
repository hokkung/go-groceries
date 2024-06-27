package user_test

import (
	"github.com/hokkung/go-groceries/internal/service/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hokkung/go-groceries/internal/repository/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite

	underTest user.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)

	s.underTest = user.NewUserService(mockUserRepository)
}

func (s *UserServiceTestSuite) TestGet() {
	one := 1
	res := s.underTest.Get(one)

	s.Equal(one, res)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
