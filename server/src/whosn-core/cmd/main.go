package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/gorilla/mux"

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

func listEvents(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "listEvents"})
	ll.Println("Endpoint Hit")
	json.NewEncoder(w).Encode(Events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "getEvent"})
	ll.Println("Endpoint Hit")

	vars := mux.Vars(r)
	eventID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Non integer event id value %v passed into handler", eventID)
		return
	}
	for _, event := range Events {
		if event.ID == uint64(eventID) {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/v1/events", listEvents)
	router.HandleFunc("/api/v1/event/{id}", getEvent)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
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
			Link:       "www.somepage.com/abasdcasdfasdf/1",
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
			Link:       "www.somepage.com/abasdcasdfasdf/2",
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
			Link:       "www.somepage.com/abasdcasdfasdf/3",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}
	handleRequests()
}
