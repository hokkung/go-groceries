package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkung/go-groceries/handler"

	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
	productHandler handler.ProductHandler
}

func (c *Customizer) Register(s *srv.Server) {
	s.Engine.GET("/ping", func(ctx *gin.Context) {
		id := c.productHandler.Get(1)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("pong %+v", id),
		})
	})
}

func NewCustomizer(productHandler handler.ProductHandler) *Customizer {
	return &Customizer{productHandler: productHandler}
}

func ProvideCustomizer(productHandler handler.ProductHandler) srv.ServerCustomizer {
	return NewCustomizer(productHandler)
}
