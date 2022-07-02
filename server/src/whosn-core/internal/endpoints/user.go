package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUser(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "GetUser", "actorID": actorID, "userID": userID})
	ll.Info("Endpoint Hit")

	if actorID != userID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to view user"})
		ctx.Abort()
		return
	}

	user, err := ds.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

func CreateUser(ctx *gin.Context) {
	ll := log.WithFields(log.Fields{"endpoint": "CreateUser"})
	ll.Info("Endpoint Hit")

	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		ll.Warn("Failed to hash password")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user.Construct()

	db := ctx.MustGet("DB").(*data.PGStore)
	err := db.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "UpdateUser", "actorID": actorID, "userID": userID})
	ll.Info("Endpoint Hit")

	if actorID != userID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to view user"})
		ctx.Abort()
		return
	}

	var userUpdate models.User
	if err := ctx.BindJSON(&userUpdate); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user, err := ds.UpdateUserByID(userUpdate, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "actorID": actorID, "userID": userID})
	ll.Info("Endpoint Hit")

	if actorID != userID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to view user"})
		ctx.Abort()
		return
	}

	err := ds.DeleteUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "{}")
}
