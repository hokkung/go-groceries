package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkung/go-groceries/internal/service/user"
	"net/http"
	"strconv"
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

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	id2, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{})
	}

	res, err := h.userService.Get(c, id2)
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}

	c.JSON(200, gin.H{"user": res})
}
