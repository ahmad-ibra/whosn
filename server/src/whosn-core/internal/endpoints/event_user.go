package endpoints

import (
	"net/http"

	wnerr "github.com/Ahmad-Ibra/whosn-core/internal/errors"
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
		if err, ok := err.(wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
			// actor hasn't joined the event
			eventUser.Construct(eventID, actorID)
			insErr := ds.InsertEventUser(*eventUser)
			if insErr != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": insErr.Error()})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, eventUser)
}

func LeaveEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "LeaveEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	// check that actor has joined the event
	eventUser, err := ds.GetEventUserByEventIDUserID(eventID, actorID)
	if err != nil {
		if err, ok := err.(wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
			// actor not in event, no need to leave
			ctx.JSON(http.StatusOK, "{}")
			ctx.Abort()
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
	}

	err = ds.DeleteEventUserByID(eventUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}
