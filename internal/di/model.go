package di

import (
	"github.com/google/wire"
	handler2 "github.com/hokkung/go-groceries/internal/handler"
	repository2 "github.com/hokkung/go-groceries/internal/repository"
	"github.com/hokkung/go-groceries/internal/server"
	service2 "github.com/hokkung/go-groceries/internal/service"
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
	repository2.ProvideGormDB,
	repository2.ProvideProductRepository,
	repository2.ProvideUserRepository,
)

var ServiceSet = wire.NewSet(
	service2.ProvideProductService,
	service2.ProvideUserService,
)

var HandlerSet = wire.NewSet(
	handler2.ProvideProductHandler,
	handler2.ProvideUserHandler,
)

type ApplicationAPI struct {
	Server *srv.Server
}
