package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppProperties struct {
	Name string `envconfig:"NAME" default:"groceries_api"`
	Env  string `envconfig:"ENV" default:"dev"`
}

// NewAppProperties creates instance
func NewAppProperties() *AppProperties {
	var app AppProperties
	err := envconfig.Process(AppPrefix, &app)
	if err != nil {
		fmt.Println("binding app configuration failed", err.Error())
		panic(err)
	}

	return &app
}

// ProvideAppProperties provides instance for di
func ProvideAppProperties() AppProperties {
	return *NewAppProperties()
}
