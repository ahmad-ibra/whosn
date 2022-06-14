package endpoints

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "Ping"})
	ll.Println("Endpoint Hit")
	fmt.Fprintf(w, "Im still alive!!!")
}
