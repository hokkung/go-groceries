package handler

import (
	"github.com/hokkung/go-groceries/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(id int) int
	Login(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func (h *userHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "hokkung",
	})
}

func (h *userHandler) Get(id int) int {
	return h.userService.Get(id)
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func ProvideUserHandler(userService service.UserService) UserHandler {
	return NewUserHandler(userService)
}
