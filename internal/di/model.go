package di

import (
	"github.com/google/wire"
	"github.com/hokkung/go-groceries/config"
	catapi "github.com/hokkung/go-groceries/internal/client/cat_api"
	ih "github.com/hokkung/go-groceries/internal/handler/item"
	ph "github.com/hokkung/go-groceries/internal/handler/product"
	uh "github.com/hokkung/go-groceries/internal/handler/user"
	"github.com/hokkung/go-groceries/internal/middleware"
	repository2 "github.com/hokkung/go-groceries/internal/repository"
	"github.com/hokkung/go-groceries/internal/server"
	"github.com/hokkung/go-groceries/internal/service/item"
	"github.com/hokkung/go-groceries/internal/service/product"
	"github.com/hokkung/go-groceries/internal/service/user"
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

	wire.Bind(new(user.UserService), new(*user.User)),
	user.ProvideUserService,

	wire.Bind(new(item.ItemService), new(*item.Item)),
	item.ProvideItem,
)

var HandlerSet = wire.NewSet(
	ph.ProvideProductHandler,
	uh.ProvideUserHandler,
	ih.ProvideItemHandler,
	middleware.ProvideGormMiddleware,
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
