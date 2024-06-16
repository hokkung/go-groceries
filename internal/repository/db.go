package repository

import (
	entity2 "github.com/hokkung/go-groceries/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideGormDB() (*gorm.DB, error) {
	return NewGormDB()
}

func NewGormDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/grocery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity2.Product{})
	db.AutoMigrate(&entity2.User{})

	return db, nil
}
