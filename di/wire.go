//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func InitializeApplication() (*ApplicationAPI, error) {
	wire.Build(APISet)
	return &ApplicationAPI{}, nil
}
