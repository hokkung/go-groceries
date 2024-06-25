package server

import (
	"fmt"
	handler2 "github.com/hokkung/go-groceries/internal/handler"
	"github.com/hokkung/go-groceries/internal/handler/item"
	"net/http"

	"github.com/gin-gonic/gin"
	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
	productHandler handler2.ProductHandler
	userHandler    handler2.UserHandler
	itemHandler    *item.ItemHandler
}

func (c *Customizer) Register(s *srv.Server) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		id := c.productHandler.Get(1)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("pong %+v", id),
		})
	})

	userGroup := s.Engine.Group("/users")
	userGroup.POST(
		"/login",
		c.userHandler.Login,
	)

	itemGroup := s.Engine.Group("/items")
	itemGroup.GET(
		"/search",
		c.itemHandler.Search,
	)
}

func NewCustomizer(
	productHandler handler2.ProductHandler,
	userHandler handler2.UserHandler,
	itemHandler *item.ItemHandler,
) *Customizer {
	return &Customizer{
		productHandler: productHandler,
		userHandler:    userHandler,
		itemHandler:    itemHandler,
	}
}

func ProvideCustomizer(
	productHandler handler2.ProductHandler,
	userHandler handler2.UserHandler,
	itemHandler *item.ItemHandler,
) (srv.ServerCustomizer, func(), error) {
	return NewCustomizer(
		productHandler,
		userHandler,
		itemHandler,
	), func() {}, nil
}
