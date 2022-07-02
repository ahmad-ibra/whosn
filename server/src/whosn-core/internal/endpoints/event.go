package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	wnerr "github.com/Ahmad-Ibra/whosn-core/internal/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListJoinedEvents(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListJoinedEvents", "actorID": actorID})
	ll.Info("Endpoint Hit")

	events, err := ds.ListJoinedEvents(actorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func ListOwnedEvents(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListOwnedEvents", "actorID": actorID})
	ll.Info("Endpoint Hit")

	events, err := ds.ListOwnedEvents(actorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "GetEvent", "actorID": actorID, "eventID": eventID})
	ll.Info("Endpoint Hit")

	event, err := ds.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, event)
	return
}

func CreateEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "CreateEvent", "actorID": actorID})
	ll.Info("Endpoint Hit")

	var event models.Event
	if err := ctx.BindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	event.Construct(actorID)

	if event.MinUsers > event.MaxUsers {
		ll.Warn("MinUsers must be less than MaxUser")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "MinUsers must be less than MaxUser"})
		ctx.Abort()
		return
	}

	err := ds.InsertEvent(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func UpdateEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "UpdateEvent", "actorID": actorID, "eventID": eventID})
	ll.Info("Endpoint Hit")

	event, err := ds.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if event.OwnerID != actorID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to update event"})
		ctx.Abort()
		return
	}

	var eventUpdate models.Event
	if err := ctx.BindJSON(&eventUpdate); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if eventUpdate.MinUsers > eventUpdate.MaxUsers {
		ll.Warn("MinUsers must be less than MaxUser")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "MinUsers must be less than MaxUser"})
		ctx.Abort()
		return
	}

	event, err = ds.UpdateEventByID(eventUpdate, eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func DeleteEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "DeleteEvent", "actorID": actorID, "eventID": eventID})
	ll.Info("Endpoint Hit")

	event, err := ds.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if event.OwnerID != actorID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to update event"})
		ctx.Abort()
		return
	}

	err = ds.DeleteEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "{}")
}

func JoinEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "JoinEvent", "actorID": actorID, "eventID": eventID})
	ll.Info("Endpoint Hit")

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
		if err, ok := err.(*wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
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
	ll.Info("Endpoint Hit")

	// check that actor has joined the event
	eventUser, err := ds.GetEventUserByEventIDUserID(eventID, actorID)
	if err != nil {
		if err, ok := err.(*wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
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
