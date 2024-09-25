package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkung/go-groceries/internal/handler"
	"github.com/hokkung/go-groceries/internal/service/user"
	"net/http"
	"strconv"
)

type Handler struct {
	userService user.UserService
}

// NewUserHandler creates instance
func NewUserHandler(userService user.UserService) *Handler {
	return &Handler{userService: userService}
}

// ProvideUserHandler provides instance for di
func ProvideUserHandler(userService user.UserService) *Handler {
	return NewUserHandler(userService)
}

// Login ...
//
//	@Summary	login
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	Helloworld
//	@Router		/users/login [post]
func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "hokkung",
	})
}

// Get ...
//
//	@Summary	get user by id
//	@Schemes
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"1"
//	@Success	200	{object}	handler.Response[User]
//	@Failure	500
//	@Router		/users/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	userIDInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userID, err := h.userService.Get(c, userIDInt)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, handler.Response[User]{
		Data: User{ID: userID},
	})
}
