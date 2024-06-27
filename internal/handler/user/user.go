package user

import (
	"github.com/hokkung/go-groceries/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService user.UserService
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "hokkung",
	})
}

// NewUserHandler creates instance
func NewUserHandler(userService user.UserService) *Handler {
	return &Handler{userService: userService}
}

// ProvideUserHandler provides instance for di
func ProvideUserHandler(userService user.UserService) *Handler {
	return NewUserHandler(userService)
}

func (h *Handler) Get(id int) int {
	return h.userService.Get(id)
}
