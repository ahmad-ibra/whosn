package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/auth"
	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var body models.CreateUserBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	user, err := ds.GetUserByUserName(body.UserName)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(body.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}
	cfg := config.GetConfig()
	tokenString, err := auth.GenerateJWT(user.ID, cfg.JWTKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
