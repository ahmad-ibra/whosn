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
	ll := log.WithFields(log.Fields{"endpoint": "GetUser", "actorID": actorID})
	ll.Info("Endpoint Hit")

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	user, err := ds.GetUserByID(actorID)
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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	err := ds.InsertUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "UpdateUser", "actorID": actorID})
	ll.Info("Endpoint Hit")

	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ll.Warn("Failed to unmarshall request body")
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

	originalUser, err := ds.GetUserByID(actorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err = user.ConstructUpdate(originalUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err = ds.UpdateUserByID(&user, actorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "actorID": actorID})
	ll.Info("Endpoint Hit")

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	err := ds.DeleteUserByID(actorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "{}")
}
