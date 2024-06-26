package di

import (
	"github.com/google/wire"
	"github.com/hokkung/go-groceries/config"
	catapi "github.com/hokkung/go-groceries/internal/client/cat_api"
	handler2 "github.com/hokkung/go-groceries/internal/handler"
	ih "github.com/hokkung/go-groceries/internal/handler/item"
	product2 "github.com/hokkung/go-groceries/internal/handler/product"
	repository2 "github.com/hokkung/go-groceries/internal/repository"
	"github.com/hokkung/go-groceries/internal/server"
	service2 "github.com/hokkung/go-groceries/internal/service"
	"github.com/hokkung/go-groceries/internal/service/item"
	"github.com/hokkung/go-groceries/internal/service/product"
	"github.com/hokkung/go-groceries/pkg/cnt"
	srv "github.com/hokkung/srv/server"
)

var APISet = wire.NewSet(
	ConfigSet,
	RepositorySet,
	ServiceSet,
	HandlerSet,
	ClientSet,
	InternalSet,
	ExternalSet,
	wire.Struct(new(ApplicationAPI), "*"),
)

var ConfigSet = wire.NewSet(
	config.ProvideAppProperties,
	config.ProvideMysqlProperties,
	config.ProvideConfiguration,
)

var RepositorySet = wire.NewSet(
	repository2.ProvideGormDB,
	repository2.ProvideProductRepository,
	repository2.ProvideUserRepository,
)

var ServiceSet = wire.NewSet(
	wire.Bind(new(product.ProductService), new(*product.Product)),
	product.ProvideProductService,
	service2.ProvideUserService,

	wire.Bind(new(item.ItemService), new(*item.Item)),
	item.ProvideItem,
)

var HandlerSet = wire.NewSet(
	product2.ProvideProductHandler,
	handler2.ProvideUserHandler,
	ih.ProvideItemHandler,
)

var ClientSet = wire.NewSet(
	wire.Bind(new(catapi.Client), new(*catapi.CatAPI)),
	catapi.ProvideCatAPI,
)

var InternalSet = wire.NewSet(
	server.ProvideCustomizer,
)

var ExternalSet = wire.NewSet(
	srv.ProvideServer,
	cnt.ProvideClient,
)

type ApplicationAPI struct {
	Server *srv.Server
	Config config.Configuration
}
