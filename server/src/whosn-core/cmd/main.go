package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"

	log "github.com/sirupsen/logrus"
)

var Events []data.Event

func homePage(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "homePage"})
	ll.Println("Endpoint Hit")

	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		ll.Warn("Failed to write a response")
		return
	}
}

func returnAllEvents(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "returnAllEvents"})
	ll.Println("Endpoint Hit")
	json.NewEncoder(w).Encode(Events)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/v1/events", returnAllEvents)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {

	Events = []data.Event{
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
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         1,
			Name:       "Soccer",
			StartTime:  time.Time{},
			Location:   "Tom binnie",
			MinUsers:   10,
			MaxUsers:   22,
			Price:      155.00,
			IsFlatRate: false,
			OwnerID:    1,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
		{
			ID:         1,
			Name:       "Movie",
			StartTime:  time.Time{},
			Location:   "Landmarks Guildford",
			MinUsers:   1,
			MaxUsers:   10,
			Price:      12,
			IsFlatRate: true,
			OwnerID:    2,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}
	handleRequests()
}
