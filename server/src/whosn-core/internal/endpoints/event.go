package endpoints

import (
	"encoding/json"
	"io/ioutil"
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
	// TODO: Break up ListEvents into ListJoinedEvents and ListOwnedEvents
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

	// TODO: return a 404
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "CreateEvent"})
	ll.Println("Endpoint Hit")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var event data.Event
	err := json.Unmarshal(reqBody, &event)
	if err != nil {
		ll.Warnf("Failed to unmarshall request body: %v", string(reqBody))
		return
	}

	curTime := time.Now()
	event.CreatedAt = curTime
	event.UpdatedAt = curTime
	event.ID = uint64(len(events)) + 1

	// TODO: fill in the OwnerID (when jwt is implemented) and generate the link.
	// This leaves the fields that should be passed in at:
	// Name, StartTime, Location, MinUsers, MaxUsers, Price, IsFlatRate

	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe update
	vars := mux.Vars(r)

	ll := log.WithFields(log.Fields{"endpoint": "UpdateUser", "userID": vars["id"]})
	ll.Println("Endpoint Hit")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var userUpdate data.User
	err := json.Unmarshal(reqBody, &userUpdate)
	if err != nil {
		ll.Warnf("Failed to unmarshall request body: %v", string(reqBody))
		return
	}

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Could not convert id %v passed into handler into an integer", vars["id"])
		return
	}

	for i := 0; i < len(users); i++ {
		user := &users[i]
		if user.ID == uint64(userID) {
			user.UpdatedAt = time.Now()
			if userUpdate.Name != "" {
				user.Name = userUpdate.Name
			}
			if userUpdate.Email != "" {
				user.Email = userUpdate.Email
			}
			if userUpdate.PhoneNumber != "" {
				user.PhoneNumber = userUpdate.PhoneNumber
			}
			return
		}
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe delete
	vars := mux.Vars(r)

	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "userID": vars["id"]})
	ll.Println("Endpoint Hit")

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Could not convert id %v passed into handler into an integer", vars["id"])
		return
	}

	for i, user := range users {
		if user.ID == uint64(userID) {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}

func JoinEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe delete
	vars := mux.Vars(r)

	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "userID": vars["id"]})
	ll.Println("Endpoint Hit")

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Could not convert id %v passed into handler into an integer", vars["id"])
		return
	}

	for i, user := range users {
		if user.ID == uint64(userID) {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}

func LeaveEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe delete
	vars := mux.Vars(r)

	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "userID": vars["id"]})
	ll.Println("Endpoint Hit")

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		ll.Warnf("Could not convert id %v passed into handler into an integer", vars["id"])
		return
	}

	for i, user := range users {
		if user.ID == uint64(userID) {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}
