package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hokkung/go-groceries/config"

	entity2 "github.com/hokkung/go-groceries/internal/entity"
)

type Entity interface {
	Table() string
}

func ProvideGormDB(config config.Configuration) (*gorm.DB, error) {
	return NewGormDB(config)
}

var entities = []Entity{
	&entity2.Product{},
	&entity2.User{},
}

func NewGormDB(config config.Configuration) (*gorm.DB, error) {
	dsn := config.MysqlProperties.DNS()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, ent := range entities {
		err = db.AutoMigrate(ent)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
