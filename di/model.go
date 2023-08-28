package di

import (
	"github.com/google/wire"
	"github.com/hokkung/go-groceries/handler"
	"github.com/hokkung/go-groceries/repository"
	"github.com/hokkung/go-groceries/server"
	"github.com/hokkung/go-groceries/service"
	srv "github.com/hokkung/srv/server"
)

var APISet = wire.NewSet(
	RepositorySet,
	ServiceSet,
	HandlerSet,
	server.ProvideCustomizer,
	srv.NewServer,
	wire.Struct(new(ApplicationAPI), "*"),
)

var RepositorySet = wire.NewSet(
	repository.ProvideGormDB,
	repository.ProvideProductRepository,
)

var ServiceSet = wire.NewSet(
	service.ProvideProductService,
)

var HandlerSet = wire.NewSet(
	handler.ProvideProductHandler,
)

type ApplicationAPI struct {
	Server *srv.Server
}
