package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// mock events till we get a db in place
	events = []data.Event{
		{
			ID:         1,
			Name:       "Volleyball",
			StartTime:  time.Time{},
			Location:   "6Pack",
			MinUsers:   10,
			MaxUsers:   12,
			Price:      120.00,
			IsFlatRate: false,
			OwnerID:    1,
			Link:       "www.somepage.com/abasdcasdfasdf/1",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         2,
			Name:       "Soccer",
			StartTime:  time.Time{},
			Location:   "Tom binnie",
			MinUsers:   10,
			MaxUsers:   22,
			Price:      155.00,
			IsFlatRate: false,
			OwnerID:    1,
			Link:       "www.somepage.com/abasdcasdfasdf/2",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         3,
			Name:       "Movie",
			StartTime:  time.Time{},
			Location:   "Landmarks Guildford",
			MinUsers:   1,
			MaxUsers:   10,
			Price:      12,
			IsFlatRate: true,
			OwnerID:    2,
			Link:       "www.somepage.com/abasdcasdfasdf/3",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}
)

func ListEvents(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "ListEvents"})
	ll.Println("Endpoint Hit")
	json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ll := log.WithFields(log.Fields{"endpoint": "GetEvent", "eventID": vars["id"]})
	ll.Println("Endpoint Hit")

	eventID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Could not convert id %v passed into handler into an integer", vars["id"])
		return
	}
	for _, event := range events {
		if event.ID == uint64(eventID) {
			json.NewEncoder(w).Encode(event)
			return
		}
	}
}
