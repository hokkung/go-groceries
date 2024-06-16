package server

import (
	"fmt"
	handler2 "github.com/hokkung/go-groceries/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
	productHandler handler2.ProductHandler
	userHandler    handler2.UserHandler
}

func (c *Customizer) Register(s *srv.Server) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		id := c.productHandler.Get(1)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("pong %+v", id),
		})
	})

	userGroup := s.Engine.Group("/users")
	userGroup.POST("/login", c.userHandler.Login)
}

func NewCustomizer(
	productHandler handler2.ProductHandler,
	userHandler handler2.UserHandler,
) *Customizer {
	return &Customizer{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func ProvideCustomizer(
	productHandler handler2.ProductHandler,
	userHandler handler2.UserHandler,
) (srv.ServerCustomizer, func(), error) {
	return NewCustomizer(
		productHandler,
		userHandler,
	), func() {}, nil
}
