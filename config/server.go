package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type ServerProperties struct {
	Port string `envconfig:"PORT" default:"8080"`
}

func NewServerProperties() *ServerProperties {
	var server ServerProperties
	err := envconfig.Process("SERVER", &server)
	if err != nil {
		fmt.Println("binding server configuration failed", err.Error())
		panic(err)
	}

	return &server
}

func ProvideServerProperties() ServerProperties {
	return *NewServerProperties()
}
