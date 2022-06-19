package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListEventUsers(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListEventUsers", "actorID": actorID})
	ll.Println("Endpoint Hit")

	eventUsers, err := ds.ListAllEventUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, eventUsers)
}

func JoinEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "JoinEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	// check that event exists
	_, err := ds.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// check that actor hasn't already joined the event
	eventUser, err := ds.GetEventUserByEventIDUserID(eventID, actorID)
	if err != nil {
		// TODO: once custom error type with status is created finish off this logic, for now assuming its NOTFOUND
		// if error is NOTFOUND {
		// construct and join the event
		eventUser.Construct(eventID, actorID)
		err = ds.InsertEventUser(*eventUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		// } else {
		// for all other error types, just return the error
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//ctx.Abort()
		//return
		// }
	}

	ctx.JSON(http.StatusOK, eventUser)
}

func LeaveEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "LeaveEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	eventUser, err := ds.GetEventUserByEventIDUserID(eventID, actorID)
	if err != nil {
		// TODO: once custom error type with status is created finish off this logic, for now assuming its NOTFOUND
		// if error is NOTFOUND {
		ctx.JSON(http.StatusOK, "{}")
		ctx.Abort()
		return
		// } else {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//ctx.Abort()
		//return
		// }
	}

	err = ds.DeleteEventUserByID(eventUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}
