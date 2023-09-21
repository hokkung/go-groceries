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
	srv.ProvideServer,
	wire.Struct(new(ApplicationAPI), "*"),
)

var RepositorySet = wire.NewSet(
	repository.ProvideGormDB,
	repository.ProvideProductRepository,
	repository.ProvideUserRepository,
)

var ServiceSet = wire.NewSet(
	service.ProvideProductService,
	service.ProvideUserService,
)

var HandlerSet = wire.NewSet(
	handler.ProvideProductHandler,
	handler.ProvideUserHandler,
)

type ApplicationAPI struct {
	Server *srv.Server
}
