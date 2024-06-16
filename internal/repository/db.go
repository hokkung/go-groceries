package repository

import (
	"github.com/hokkung/go-groceries/config"
	entity2 "github.com/hokkung/go-groceries/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideGormDB(config config.Configuration) (*gorm.DB, error) {
	return NewGormDB(config)
}

func NewGormDB(config config.Configuration) (*gorm.DB, error) {
	dsn := config.MysqlProperties.DNS()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity2.Product{})
	db.AutoMigrate(&entity2.User{})

	return db, nil
}
