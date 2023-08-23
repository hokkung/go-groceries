package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/grocery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
