//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/google/wire"
)

func InitializeApplication(context context.Context) (*ApplicationAPI, func(), error) {
	wire.Build(APISet)
	return &ApplicationAPI{}, func() {}, nil
}
