package main

import (
	"net/http"
	"os"

	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/_hc", endpoints.Ping)

	router.HandleFunc("/api/v1/user/{id}", endpoints.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/user/{id}", endpoints.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", endpoints.GetUser)
	router.HandleFunc("/api/v1/user", endpoints.CreateUser).Methods("POST")

	router.HandleFunc("/api/v1/events", endpoints.ListEvents)
	router.HandleFunc("/api/v1/event/{id}", endpoints.DeleteEvent).Methods("DELETE")
	router.HandleFunc("/api/v1/event/{id}", endpoints.UpdateEvent).Methods("PUT")
	router.HandleFunc("/api/v1/event/{id}", endpoints.GetEvent)
	router.HandleFunc("/api/v1/event/{id}/join", endpoints.JoinEvent)
	router.HandleFunc("/api/v1/event/{id}/leave", endpoints.LeaveEvent)
	router.HandleFunc("/api/v1/event", endpoints.CreateEvent).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	handleRequests()
}
