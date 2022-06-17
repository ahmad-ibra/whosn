package endpoints

import (
	"net/http"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	// mock events till we get a db in place
	events = []models.Event{
		{
			ID:         "f503857c-5334-450d-be87-15bdcde50341",
			Name:       "Volleyball",
			StartTime:  time.Time{},
			Location:   "6Pack",
			MinUsers:   10,
			MaxUsers:   12,
			Price:      120.00,
			IsFlatRate: false,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50342",
			Link:       "www.somepage.com/abasdcasdfasdf/1",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         "50262b10-3d8e-4134-9869-1e0ed5cfe9f7",
			Name:       "Soccer",
			StartTime:  time.Time{},
			Location:   "Tom binnie",
			MinUsers:   10,
			MaxUsers:   22,
			Price:      155.00,
			IsFlatRate: false,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50343",
			Link:       "www.somepage.com/abasdcasdfasdf/2",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         "45de396a-4880-4c52-9689-f8812bf67a51",
			Name:       "Movie",
			StartTime:  time.Time{},
			Location:   "Landmarks Guildford",
			MinUsers:   1,
			MaxUsers:   10,
			Price:      12,
			IsFlatRate: true,
			OwnerID:    "f503857c-5334-450d-be87-15bdcde50344",
			Link:       "www.somepage.com/abasdcasdfasdf/3",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}
)

func ListEvents(ctx *gin.Context) {
	// TODO: Break up ListEvents into ListJoinedEvents and ListOwnedEvents
	actorID := ctx.GetString("actorID")
	ll := log.WithFields(log.Fields{"endpoint": "ListEvents", "actorID": actorID})
	ll.Println("Endpoint Hit")
	ctx.JSON(http.StatusOK, events)
}

func GetEvent(ctx *gin.Context) {
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "GetEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	for _, event := range events {
		if event.ID == eventID {
			ctx.JSON(http.StatusOK, event)
			return
		}
	}

	ll.Warn("Event not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	ctx.Abort()
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

	events = append(events, event)
	ctx.JSON(http.StatusOK, event)
}

func UpdateEvent(ctx *gin.Context) {
	// TODO: make this a thread safe update
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "UpdateEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	var eventUpdate models.Event
	if err := ctx.BindJSON(&eventUpdate); err != nil {
		ll.Warn("Failed to unmarshall request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	for i := 0; i < len(events); i++ {
		event := &events[i]
		if event.ID == eventID {
			if event.OwnerID != actorID {
				ll.Warn("Unauthorized")
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to update event"})
				ctx.Abort()
				return
			}
			event.UpdatedAt = time.Now()
			if eventUpdate.Name != "" {
				event.Name = eventUpdate.Name
			}
			if !eventUpdate.StartTime.IsZero() {
				event.StartTime = eventUpdate.StartTime
			}
			if eventUpdate.Location != "" {
				event.Location = eventUpdate.Location
			}
			// TODO: dont allow setting MinUsers above MaxUsers
			if eventUpdate.MinUsers != 0 {
				event.MinUsers = eventUpdate.MinUsers
			}
			if eventUpdate.MaxUsers != 0 {
				event.MaxUsers = eventUpdate.MaxUsers
			}
			if eventUpdate.Price != 0 {
				event.Price = eventUpdate.Price
			}
			// Note: frontend needs to make sure that its always passing this value through
			if eventUpdate.IsFlatRate != event.IsFlatRate {
				event.IsFlatRate = eventUpdate.IsFlatRate
			}
			ctx.JSON(http.StatusOK, event)
			return
		}
	}
	ll.Warn("Event not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	ctx.Abort()
}

func DeleteEvent(ctx *gin.Context) {
	// TODO: make this a thread safe delete
	actorID := ctx.GetString("actorID")
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "DeleteEvent", "actorID": actorID, "eventID": eventID})
	ll.Println("Endpoint Hit")

	for i, event := range events {
		if event.ID == eventID {
			if event.OwnerID != actorID {
				ll.Warn("Unauthorized")
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Actor not authorized to delete event"})
				ctx.Abort()
				return
			}
			events = append(events[:i], events[i+1:]...)
			ctx.JSON(http.StatusOK, "{}")
			return
		}
	}
	ll.Warn("Event not found")
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	ctx.Abort()
}

func JoinEvent(ctx *gin.Context) {
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "JoinEvent", "eventID": eventID})
	ll.Println("Endpoint Hit")

	ll.Print("TODO: implement")
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}

func LeaveEvent(ctx *gin.Context) {
	eventID := ctx.Param("id")
	ll := log.WithFields(log.Fields{"endpoint": "LeaveEvent", "eventID": eventID})
	ll.Println("Endpoint Hit")

	ll.Print("TODO: implement")
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}
