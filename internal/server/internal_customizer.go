package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	srv "github.com/hokkung/srv/server"
)

type InternalCustomizer struct {
}

func NewInternalCustomizer() *InternalCustomizer {
	return &InternalCustomizer{}
}

func ProvideInternalCustomizer() srv.ServerCustomizer {
	return NewInternalCustomizer()
}

func (c *InternalCustomizer) Register(s *srv.Server) {
	s.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
