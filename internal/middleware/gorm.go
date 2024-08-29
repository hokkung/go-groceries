package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormMiddleware struct {
	db *gorm.DB
}

func NewGormMiddleware(db *gorm.DB) *GormMiddleware {
	return &GormMiddleware{
		db: db,
	}
}

func ProvideGormMiddleware(db *gorm.DB) *GormMiddleware {
	return NewGormMiddleware(db)
}

func (mw *GormMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := mw.db.Begin()
		c.Set("db", tx)
		c.Next()

		if len(c.Errors) <= 0 {
			if err := tx.Commit(); err != nil {
				tx.Rollback()
			}
			return
		}

		tx.Rollback()
	}
}
