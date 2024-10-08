// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/hokkung/go-groceries/config"
	"github.com/hokkung/go-groceries/internal/client/cat_api"
	item2 "github.com/hokkung/go-groceries/internal/handler/item"
	product2 "github.com/hokkung/go-groceries/internal/handler/product"
	user2 "github.com/hokkung/go-groceries/internal/handler/user"
	"github.com/hokkung/go-groceries/internal/middleware"
	"github.com/hokkung/go-groceries/internal/repository"
	"github.com/hokkung/go-groceries/internal/server"
	"github.com/hokkung/go-groceries/internal/service/item"
	"github.com/hokkung/go-groceries/internal/service/product"
	"github.com/hokkung/go-groceries/internal/service/user"
	server2 "github.com/hokkung/srv/server"
)

// Injectors from wire.go:

func InitializeApplication(context2 context.Context) (*ApplicationAPI, func(), error) {
	mysqlProperties := config.ProvideMysqlProperties()
	db, err := repository.ProvideGormDB(mysqlProperties)
	if err != nil {
		return nil, nil, err
	}
	gormMiddleware := middleware.ProvideGormMiddleware(db)
	productRepository, cleanup, err := repository.ProvideProductRepository(db)
	if err != nil {
		return nil, nil, err
	}
	productProduct := product.ProvideProductService(productRepository)
	product3 := product2.ProvideProductHandler(productProduct)
	userRepository, cleanup2, err := repository.ProvideUserRepository(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userUser := user.ProvideUserService(userRepository)
	handler := user2.ProvideUserHandler(userUser)
	catAPI, err := catapi.ProvideCatAPI()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	itemItem := item.ProvideItem(catAPI)
	itemHandler := item2.ProvideItemHandler(itemItem)
	serverCustomizer, cleanup3, err := server.ProvideCustomizer(gormMiddleware, product3, handler, itemHandler)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serverServer, cleanup4, err := server2.ProvideServer(serverCustomizer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	appProperties := config.ProvideAppProperties()
	configuration := config.ProvideConfiguration(appProperties, mysqlProperties)
	applicationAPI := &ApplicationAPI{
		Server: serverServer,
		Config: configuration,
	}
	return applicationAPI, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
