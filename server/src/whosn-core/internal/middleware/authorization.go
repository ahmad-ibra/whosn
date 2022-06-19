package middleware

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/auth"
	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data"

	"github.com/gin-gonic/gin"
)

var ds = data.GetInMemoryStore()

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		cfg := config.GetConfig()
		actorID, err := auth.ValidateToken(tokenString, cfg.JWTKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		_, err = ds.GetUserByID(actorID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("actorID", actorID)
		ctx.Next()
	}
}
