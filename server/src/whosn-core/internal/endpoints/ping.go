package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Ping(ctx *gin.Context) {
	ll := log.WithFields(log.Fields{"endpoint": "Ping"})
	ll.Println("Endpoint Hit")

	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
