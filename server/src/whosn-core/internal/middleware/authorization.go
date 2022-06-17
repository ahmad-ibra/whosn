package middleware

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		actorID, err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("actorID", actorID)
		ctx.Next()
	}
}