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

var Groups []data.Group

func homePage(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "homePage"})
	ll.Println("Endpoint Hit")

	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		ll.Warn("Failed to write a response")
		return
	}
}

func returnAllGroups(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "returnAllGroups"})
	ll.Println("Endpoint Hit")
	json.NewEncoder(w).Encode(Groups)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/v1/groups", returnAllGroups)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	Groups = []data.Group{
		{
			ID:        1,
			Name:      "bobcorn",
			OwnerID:   1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			ID:        2,
			Name:      "balbooliches",
			OwnerID:   1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			ID:        3,
			Name:      "jiggers",
			OwnerID:   2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	handleRequests()
}
