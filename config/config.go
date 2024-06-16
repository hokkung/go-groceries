package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

const AppPrefix = "APP"

type Configuration struct {
	AppProperties   AppProperties
	MysqlProperties MysqlProperties
}

type AppProperties struct {
	Name string `envconfig:"NAME" default:"groceries_api"`
}

type MysqlProperties struct {
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     string `envconfig:"MYSQL_PORT" default:"3306"`
	Name     string `envconfig:"MYSQL_DB_NAME" default:"grocery"`
	UserName string `envconfig:"MYSQL_USERNAME" default:"root"`
	Password string `envconfig:"MYSQL_PASSWORD" default:"root"`
}

func (c *MysqlProperties) DNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

// NewConfiguration returns new config instance
func NewConfiguration() *Configuration {
	var app AppProperties
	err := envconfig.Process(AppPrefix, &app)
	if err != nil {
		fmt.Println("binding app configuration failed", err.Error())
		panic(err)
	}

	var mysql MysqlProperties
	err = envconfig.Process(AppPrefix, &mysql)
	if err != nil {
		fmt.Println("binding mysql configuration failed", err.Error())
		panic(err)
	}

	return &Configuration{
		AppProperties:   app,
		MysqlProperties: mysql,
	}
}

// ProvideConfiguration provides configuration for di
func ProvideConfiguration() Configuration {
	return *NewConfiguration()
}
