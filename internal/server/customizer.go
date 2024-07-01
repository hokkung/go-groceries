package server

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hokkung/go-groceries/graph"
	"github.com/hokkung/go-groceries/internal/handler/item"
	ph "github.com/hokkung/go-groceries/internal/handler/product"
	"github.com/hokkung/go-groceries/internal/handler/user"

	"net/http"

	"github.com/gin-gonic/gin"
	srv "github.com/hokkung/srv/server"
)

type Customizer struct {
	productHandler *ph.Product
	userHandler    *user.Handler
	itemHandler    *item.ItemHandler
}

// NewCustomizer creates instance
func NewCustomizer(
	productHandler *ph.Product,
	userHandler *user.Handler,
	itemHandler *item.ItemHandler,
) *Customizer {
	return &Customizer{
		productHandler: productHandler,
		userHandler:    userHandler,
		itemHandler:    itemHandler,
	}
}

// ProvideCustomizer provides instance for di
func ProvideCustomizer(
	productHandler *ph.Product,
	userHandler *user.Handler,
	itemHandler *item.ItemHandler,
) (srv.ServerCustomizer, func(), error) {
	return NewCustomizer(
		productHandler,
		userHandler,
		itemHandler,
	), func() {}, nil
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

	// TODO: refactor this
	gqlSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	s.Engine.POST("/query", func(c *gin.Context) {
		gqlSrv.ServeHTTP(c.Writer, c.Request)
	})

	// TODO: refactor this
	gqlPgh := playground.Handler("GraphQL playground", "/query")
	s.Engine.GET("/", func(c *gin.Context) {
		gqlPgh.ServeHTTP(c.Writer, c.Request)
	})
}
