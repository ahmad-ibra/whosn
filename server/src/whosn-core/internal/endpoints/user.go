package endpoints

import (
	"net/http"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/models"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	// mock users till we get a db in place
	users = []models.User{
		{
			ID:          "7076f342-fd08-4d44-a7ca-baeb31e581fe",
			Name:        "Ahmad I",
			Email:       "email1@whosn.xyz.com",
			PhoneNumber: "604-534-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "b1be816f-fb34-4ab4-a1de-d3a08eca5217",
			Name:        "Karrar A",
			Email:       "email23234234@whosn.xyz.com",
			PhoneNumber: "778-111-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "489c800e-034b-4225-bfb1-3327652b63cb",
			Name:        "Wael A",
			Email:       "anotherEmail@whosn.xyz.com",
			PhoneNumber: "123-345-4567",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}
)

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "GetUser", "userID": userID})
	ll.Println("Endpoint Hit")

	for _, user := range users {
		if user.ID == userID {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}

	ll.Warn("User not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func CreateUser(ctx *gin.Context) {
	ll := log.WithFields(log.Fields{"endpoint": "CreateUser"})
	ll.Println("Endpoint Hit")

	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to unmarshall request body"})
		return
	}

	curTime := time.Now()
	user.CreatedAt = curTime
	user.UpdatedAt = curTime
	user.ID = uuid.New().String()

	users = append(users, user)
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	// TODO: make this a thread safe update
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "UpdateUser", "userID": userID})
	ll.Println("Endpoint Hit")

	var userUpdate models.User
	if err := ctx.BindJSON(&userUpdate); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to unmarshall request body"})
		return
	}

	for i := 0; i < len(users); i++ {
		user := &users[i]
		if user.ID == userID {
			user.UpdatedAt = time.Now()
			if userUpdate.Name != "" {
				user.Name = userUpdate.Name
			}
			if userUpdate.Email != "" {
				user.Email = userUpdate.Email
			}
			if userUpdate.PhoneNumber != "" {
				user.PhoneNumber = userUpdate.PhoneNumber
			}
			if userUpdate.Username != "" {
				user.Username = userUpdate.Username
			}
			if userUpdate.Password != "" {
				user.Password = userUpdate.Password
			}
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
	ll.Warn("User not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func DeleteUser(ctx *gin.Context) {
	// TODO: make this a thread safe delete
	userID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "userID": userID})
	ll.Println("Endpoint Hit")

	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			ctx.JSON(http.StatusOK, "{}")
			return
		}
	}
	ll.Warn("User not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
