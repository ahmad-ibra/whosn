package middleware

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/gin-gonic/gin"
)

func DB(db *data.PGStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("DB", db)
		ctx.Next()
	}
}
