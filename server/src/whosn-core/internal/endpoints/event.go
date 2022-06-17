package endpoints

import (
	"net/http"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ListEvents(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListEvents", "actorID": actorID})
	ll.Println("Endpoint Hit")

	events, err := ds.ListAllEvents()
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
	ll.Println("Endpoint Hit")

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
	ll.Println("Endpoint Hit")

	var event models.Event
	if err := ctx.BindJSON(&event); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	event.Construct(actorID)

	if event.MinUsers > event.MaxUsers {
		ll.Warn("MinUsers must be less than MaxUser")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "MinUsers must be less than MaxUser"})
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
	ll.Println("Endpoint Hit")

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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "MinUsers must be less than MaxUser"})
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
	ll.Println("Endpoint Hit")

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
	ll.Println("Endpoint Hit")

	// - GetEventUserByEventIDUserID
	// - if no result, construct EventUser and add it (InsertEventUser), then return it
	// - else, return found UserEvent -> ctx.JSON(http.StatusOK, userEvent)
}

func LeaveEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "LeaveEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	// - GetEventUserByEventIDUserID
	// - if found result, remove the UserEvent (DeleteEventUserByEventIDUserID)
	// - else, return -> ctx.JSON(http.StatusOK, "{}")

}
