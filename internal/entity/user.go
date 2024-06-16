package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}
