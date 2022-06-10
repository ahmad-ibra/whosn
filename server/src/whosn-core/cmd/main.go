package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "homePage"})
	ll.Println("Endpoint Hit")

	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		ll.Warn("Failed to write a response")
		return
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequests()
}
