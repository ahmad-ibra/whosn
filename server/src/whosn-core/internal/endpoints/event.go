package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	wnerr "github.com/Ahmad-Ibra/whosn-core/internal/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListJoinedEvents(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListJoinedEvents", "actorID": actorID})
	ll.Info("Endpoint Hit")

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

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
	event.ID = actorID

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	err := ds.InsertEvent(&event)
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

	var event models.Event
	if err := ctx.BindJSON(&event); err != nil {
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

	originalEvent, err := ds.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if originalEvent.OwnerID != actorID {
		ll.Warn("Unauthorized")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to update event"})
		ctx.Abort()
		return
	}

	event.ConstructUpdate(originalEvent)

	err = ds.UpdateEventByID(&event, eventID)
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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

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
			eventUser = &models.EventUser{}
			eventUser.ConstructCreate(eventID, actorID)
			insErr := ds.InsertEventUser(eventUser)
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

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	// check that actor has joined the event
	eventUser, err := ds.GetEventUserByEventIDUserID(eventID, actorID)
	if err != nil {
		if err, ok := err.(*wnerr.WnError); ok && err.StatusCode == http.StatusNotFound {
			// actor not in event, no need to leave
			ctx.JSON(http.StatusOK, "{}")
			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
	}

	err = ds.DeleteEventUserByEventIDUserID(eventUser.EventID, eventUser.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, "{}")
}

func ListEventUsers(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "ListEventUsers", "actorID": actorID, "eventID": eventID})
	ll.Info("Endpoint Hit")

	ds, ok := ctx.Value("DB").(*data.PGStore)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get database from context"})
		ctx.Abort()
		return
	}

	eventUsersIn, err := ds.ListEventUsers(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, eventUsersIn)
}
