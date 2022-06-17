package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/auth"
	"github.com/Ahmad-Ibra/whosn-core/internal/models"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var request TokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	var curUser models.User
	foundUser := false

	// get user with same username
	for _, user := range users {
		if user.Username == request.Username {
			curUser = user
			foundUser = true
		}
	}

	if !foundUser {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}
	credentialError := curUser.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(curUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
