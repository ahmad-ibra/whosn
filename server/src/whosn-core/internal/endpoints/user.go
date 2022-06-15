package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// mock users till we get a db in place
	users = []data.User{
		{
			ID:          "7076f342-fd08-4d44-a7ca-baeb31e581fe",
			Name:        "Ahmad I",
			Email:       "email1@whosn.xyz.com",
			PhoneNumber: "604-534-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "b1be816f-fb34-4ab4-a1de-d3a08eca5217",
			Name:        "Karrar A",
			Email:       "email23234234@whosn.xyz.com",
			PhoneNumber: "778-111-6333",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          "489c800e-034b-4225-bfb1-3327652b63cb",
			Name:        "Wael A",
			Email:       "anotherEmail@whosn.xyz.com",
			PhoneNumber: "123-345-4567",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	ll := log.WithFields(log.Fields{"endpoint": "GetUser", "userID": userID})
	ll.Println("Endpoint Hit")

	for _, user := range users {
		if user.ID == userID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// TODO: return a 404
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ll := log.WithFields(log.Fields{"endpoint": "CreateUser"})
	ll.Println("Endpoint Hit")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var user data.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		ll.Warnf("Failed to unmarshall request body: %v", string(reqBody))
		return
	}

	curTime := time.Now()
	user.CreatedAt = curTime
	user.UpdatedAt = curTime
	user.ID = uuid.New().String()

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe update
	vars := mux.Vars(r)
	userID := vars["id"]

	ll := log.WithFields(log.Fields{"endpoint": "UpdateUser", "userID": userID})
	ll.Println("Endpoint Hit")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var userUpdate data.User
	err := json.Unmarshal(reqBody, &userUpdate)
	if err != nil {
		ll.Warnf("Failed to unmarshall request body: %v", string(reqBody))
		return
	}

	for i := 0; i < len(users); i++ {
		user := &users[i]
		if user.ID == userID {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: make this a thread safe delete
	vars := mux.Vars(r)
	userID := vars["id"]

	ll := log.WithFields(log.Fields{"endpoint": "DeleteUser", "userID": userID})
	ll.Println("Endpoint Hit")

	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}
