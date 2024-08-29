package repository

import (
	"context"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func New(db *gorm.DB) {
	gormDB = db
}

func Txn(ctx context.Context, fn func(tx *gorm.DB) error) error {
	val, ok := ctx.Value(GormContext).(*gorm.DB)
	if !ok {
		return fn(gormDB)
	}
	return fn(val)
}
