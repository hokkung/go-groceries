package repository

import (
	"context"
	"gorm.io/gorm"
)

type Base struct {
	db *gorm.DB
}

func NewBase(db *gorm.DB) *Base {
	return &Base{
		db: db,
	}
}

func (b *Base) getDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(GormContext).(*gorm.DB)
	if !ok {
		return b.db
	}
	return tx
}
