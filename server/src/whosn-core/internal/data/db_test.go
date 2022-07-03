package data

import (
	"os"
	"testing"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("ENV", "test")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "dev")
	os.Setenv("POSTGRES_PASSWORD", "pass")
	os.Setenv("POSTGRES_DBNAME", "whosn")

	db, _ := NewDB()
	cleanTables(db.Conn)
}

func cleanTables(db *pg.DB) {
	var eventUsers []models.EventUser
	db.Model(&eventUsers).Select()
	db.Model(&eventUsers).Where("event_id IS NOT NULL").Delete()

	var events []models.Event
	db.Model(&events).Select()
	db.Model(&events).Where("id IS NOT NULL").Delete()

	var users []models.User
	db.Model(&users).Select()
	db.Model(&users).Where("id IS NOT NULL").Delete()
}

func TestInsertUser(t *testing.T) {
	duplicateUser := &models.User{
		Name:        "dup name",
		UserName:    "dupUserName",
		Password:    "password",
		Email:       "dupEmail@foo.bar",
		PhoneNumber: "604-155-5555",
	}

	user := &models.User{
		Name:        "TestInsertUser Name",
		UserName:    "TestInsertUser Username",
		Password:    "TestInsertUserPassword",
		Email:       "TestInsertUser@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	var tests = []struct {
		title string
		user  *models.User
		fail  bool
	}{
		{
			title: "fails to insert user with no Name",
			user: &models.User{
				UserName:    user.UserName,
				Password:    user.Password,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
			},
			fail: true,
		},
		{
			title: "fails to insert user with no UserName",
			user: &models.User{
				Name:        user.Name,
				Password:    user.Password,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
			},
			fail: true,
		},
		{
			title: "fails to insert user with no Password",
			user: &models.User{
				Name:        user.Name,
				UserName:    user.UserName,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
			},
			fail: true,
		},
		{
			title: "fails to insert user with no Email",
			user: &models.User{
				Name:        user.Name,
				UserName:    user.UserName,
				Password:    user.Password,
				PhoneNumber: user.PhoneNumber,
			},
			fail: true,
		},
		{
			title: "fails to insert user with no PhoneNumber",
			user: &models.User{
				Name:     user.Name,
				UserName: user.UserName,
				Password: user.Password,
				Email:    user.Email,
			},
			fail: true,
		},
		{
			title: "fails to inserts a duplicate user",
			user:  duplicateUser,
			fail:  true,
		},
		{
			title: "successfully inserts a user",
			user:  user,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(duplicateUser).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.InsertUser(tt.user)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetUserByUserName(t *testing.T) {
	user := &models.User{
		Name:        "TestGetUserByUserName Name",
		UserName:    "TestGetUserByUserName-UserName",
		Password:    "TestGetUserByUserNamePassword",
		Email:       "TestGetUserByUserName@foo.bar",
		PhoneNumber: "604-555-5551",
	}

	var tests = []struct {
		title    string
		userName string
		fail     bool
	}{
		{
			title:    "fails to find user not in db",
			userName: "notInDB",
			fail:     true,
		},
		{
			title:    "successfully finds user in db",
			userName: user.UserName,
			fail:     false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			foundUser, err := db.GetUserByUserName(tt.userName)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.userName, foundUser.UserName)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	userID := uuid.New().String()
	user := &models.User{
		ID:          userID,
		Name:        "TestGetUserByID name",
		UserName:    "TestGetUserByID-username",
		Password:    "TestGetUserByIDPassword",
		Email:       "TestGetUserByID@foo.bar",
		PhoneNumber: "604-555-5533",
	}

	var tests = []struct {
		title string
		id    string
		fail  bool
	}{
		{
			title: "fails to find user if ID is not a uuid",
			id:    "notInDB",
			fail:  true,
		},
		{
			title: "fails to find user if ID is not in db",
			id:    uuid.New().String(),
			fail:  true,
		},
		{
			title: "successfully finds user in db",
			id:    user.ID,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			foundUser, err := db.GetUserByID(tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.id, foundUser.ID)
			}
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	userID := uuid.New().String()
	user := &models.User{
		ID:          userID,
		Name:        "TestDeleteUserByID name",
		UserName:    "TestDeleteUserByID-username",
		Password:    "TestDeleteUserByIDPassword",
		Email:       "TestDeleteUserByID@foo.bar",
		PhoneNumber: "604-552-5555",
	}

	var tests = []struct {
		title string
		id    string
		fail  bool
	}{
		{
			title: "fails to delete user if ID is not a uuid",
			id:    "notInDB",
			fail:  true,
		},
		{
			title: "returns no error if ID is not in db",
			id:    uuid.New().String(),
			fail:  false,
		},
		{
			title: "successfully deletes user in db",
			id:    user.ID,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.DeleteUserByID(tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateUserByID(t *testing.T) {
	createTime := time.Now().UTC()
	updateTime := time.Now().UTC().Add(time.Hour * time.Duration(5))
	userID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestUpdateUserByID name",
		UserName:    "TestUpdateUserByID-username",
		Password:    "TestUpdateUserByIDPassword",
		Email:       "TestUpdateUserByID@foo.bar",
		PhoneNumber: "624-555-5555",
		CreatedAt:   createTime,
		UpdatedAt:   createTime,
	}

	updatedUser := &models.User{
		ID:          userID,
		Name:        "TestUpdateUserByID Newname",
		UserName:    "TestUpdateUserByID-NewUsername",
		Password:    "TestUpdateUserByIDNewPassword",
		Email:       "TestUpdateUserByIDNew@foo.bar",
		PhoneNumber: "622-555-5555",
		CreatedAt:   createTime,
		UpdatedAt:   updateTime,
	}

	var tests = []struct {
		title   string
		id      string
		expUser *models.User
		fail    bool
	}{
		{
			title:   "fails to update user if ID is not a uuid",
			id:      "notInDB",
			expUser: user,
			fail:    true,
		},
		{
			title:   "returns no error if ID is not in db",
			id:      uuid.New().String(),
			expUser: user,
			fail:    false,
		},
		{
			title:   "successfully updates user in db",
			id:      user.ID,
			expUser: updatedUser,
			fail:    false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.UpdateUserByID(updatedUser, tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			upUser, _ := db.GetUserByID(user.ID)
			assert.Equal(t, tt.expUser, upUser)
		})
	}
}

func TestInsertEvent(t *testing.T) {
	userID := uuid.New().String()
	user := &models.User{
		ID:          userID,
		Name:        "TestInsertEvent name",
		UserName:    "TestInsertEvent-username",
		Password:    "TestInsertEventpassword",
		Email:       "TestInsertEvent@foo.bar",
		PhoneNumber: "602-555-5555",
	}

	event := &models.Event{
		Name:     "TestInsertEvent eventname",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestInsertEvent over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEvent.com",
	}

	var tests = []struct {
		title string
		event *models.Event
		fail  bool
	}{
		{
			title: "fails to insert event with no Name",
			event: &models.Event{
				OwnerID:  user.ID,
				Time:     event.Time,
				Location: event.Location,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Price:    event.Price,
				Link:     event.Link,
			},
			fail: true,
		},
		{
			title: "fails to insert event with no OwnerID",
			event: &models.Event{
				Name:     event.Name,
				Time:     event.Time,
				Location: event.Location,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Price:    event.Price,
				Link:     event.Link,
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Time",
			event: &models.Event{
				Name:     event.Name,
				OwnerID:  user.ID,
				Location: event.Location,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Price:    event.Price,
				Link:     event.Link,
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Location",
			event: &models.Event{
				Name:     event.Name,
				OwnerID:  user.ID,
				Time:     event.Time,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Price:    event.Price,
				Link:     event.Link,
			},
			fail: true,
		},
		{
			title: "fails to insert event with MinUsers > MaxUsers",
			event: &models.Event{
				Name:     event.Name,
				OwnerID:  user.ID,
				Location: event.Location,
				Time:     event.Time,
				MinUsers: 10,
				MaxUsers: 4,
				Price:    event.Price,
				Link:     event.Link,
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Link",
			event: &models.Event{
				Name:     event.Name,
				OwnerID:  user.ID,
				Location: event.Location,
				Time:     event.Time,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Price:    event.Price,
			},
			fail: true,
		},
		{
			title: "successfully inserts event with no Price",
			event: &models.Event{
				Name:     event.Name,
				OwnerID:  user.ID,
				Location: event.Location,
				Time:     event.Time,
				MinUsers: event.MinUsers,
				MaxUsers: event.MaxUsers,
				Link:     event.Link,
			},
			fail: false,
		},
		{
			title: "successfully inserts an event",
			event: event,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.InsertEvent(tt.event)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetEventByID(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestGetEventByID name",
		UserName:    "TestGetEventByID-username",
		Password:    "TestGetEventByIDpassword",
		Email:       "TestGetEventByID@foo.bar",
		PhoneNumber: "604-555-2255",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "TestGetEventByID- event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestGetEventByID - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestGetEventByID.com",
	}

	var tests = []struct {
		title string
		id    string
		fail  bool
	}{
		{
			title: "fails to find event if ID is not a uuid",
			id:    "notInDB",
			fail:  true,
		},
		{
			title: "fails to find event if ID is not in db",
			id:    uuid.New().String(),
			fail:  true,
		},
		{
			title: "successfully finds event in db",
			id:    event.ID,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			foundEvent, err := db.GetEventByID(tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.id, foundEvent.ID)
			}
		})
	}
}

func TestDeleteEventByID(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestDeleteEventByID name",
		UserName:    "TestDeleteEventByID-username",
		Password:    "TestDeleteEventByIDpassword",
		Email:       "TestDeleteEventByID@foo.bar",
		PhoneNumber: "604-123-5555",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "TestDeleteEventByID event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestDeleteEventByID - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestDeleteEventByID.com",
	}

	var tests = []struct {
		title string
		id    string
		fail  bool
	}{
		{
			title: "fails to delete event if ID is not a uuid",
			id:    "notInDB",
			fail:  true,
		},
		{
			title: "returns no error if ID is not in db",
			id:    uuid.New().String(),
			fail:  false,
		},
		{
			title: "successfully deletes event in db",
			id:    event.ID,
			fail:  false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.DeleteEventByID(tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdateEventByID(t *testing.T) {
	createTime := time.Now().UTC()
	updateTime := time.Now().UTC().Add(time.Hour * time.Duration(5))
	userID := uuid.New().String()
	eventID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestUpdateEventByID name",
		UserName:    "TestUpdateEventByID-username",
		Password:    "TestUpdateEventByIDpassword",
		Email:       "TestUpdateEventByID@foo.bar",
		PhoneNumber: "604-555-2345",
	}

	event := &models.Event{
		ID:        eventID,
		Name:      "TestUpdateEventByID- event name",
		OwnerID:   user.ID,
		Time:      createTime,
		Location:  "TestUpdateEventByID - over there!",
		MinUsers:  1,
		MaxUsers:  4,
		Price:     10.23,
		Link:      "http://TestUpdateEventByID.com",
		CreatedAt: createTime,
		UpdatedAt: createTime,
	}

	updatedEvent := &models.Event{
		ID:        eventID,
		Name:      "TestUpdateEventByID new name",
		OwnerID:   user.ID,
		Time:      createTime,
		Location:  "TestUpdateEventByID new location over there!",
		MinUsers:  3,
		MaxUsers:  7,
		Price:     12.23,
		Link:      "http://TestUpdateEventByID.com",
		CreatedAt: createTime,
		UpdatedAt: updateTime,
	}

	var tests = []struct {
		title    string
		id       string
		expEvent *models.Event
		fail     bool
	}{
		{
			title:    "fails to update user if ID is not a uuid",
			id:       "notInDB",
			expEvent: event,
			fail:     true,
		},
		{
			title:    "returns no error if ID is not in db",
			id:       uuid.New().String(),
			expEvent: event,
			fail:     false,
		},
		{
			title:    "successfully updates user in db",
			id:       event.ID,
			expEvent: updatedEvent,
			fail:     false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.UpdateEventByID(updatedEvent, tt.id)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			upEvent, _ := db.GetEventByID(event.ID)
			assert.Equal(t, tt.expEvent, upEvent)
		})
	}
}

func TestGetEventUserByEventIDUserID(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()
	user := &models.User{
		ID:          userID,
		Name:        "TestGetEventUserByEventIDUserID name",
		UserName:    "TestGetEventUserByEventIDUserID-username",
		Password:    "TestGetEventUserByEventIDUserIDpassword",
		Email:       "TestGetEventUserByEventIDUserID@foo.bar",
		PhoneNumber: "604-987-5555",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "TestGetEventUserByEventIDUserID- event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestGetEventUserByEventIDUserID -over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestGetEventUserByEventIDUserID.com",
	}

	eventUser := &models.EventUser{
		EventID: event.ID,
		UserID:  user.ID,
	}

	var tests = []struct {
		title   string
		eventID string
		userID  string
		fail    bool
	}{
		{
			title:   "fails to find event_user if userID is not a uuid",
			eventID: "notInDB",
			userID:  user.ID,
			fail:    true,
		},
		{
			title:   "fails to find event_user if eventID is not a uuid",
			eventID: event.ID,
			userID:  "notInDB",
			fail:    true,
		},
		{
			title:   "fails to find event_user if eventID is not in db",
			eventID: uuid.New().String(),
			userID:  user.ID,
			fail:    true,
		},
		{
			title:   "fails to find event_user if userID is not in db",
			eventID: event.ID,
			userID:  uuid.New().String(),
			fail:    true,
		},
		{
			title:   "successfully finds user in db",
			eventID: event.ID,
			userID:  user.ID,
			fail:    false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()
	db.Conn.Model(eventUser).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			foundEventUser, err := db.GetEventUserByEventIDUserID(tt.eventID, tt.userID)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.userID, foundEventUser.UserID)
				assert.Equal(t, tt.eventID, foundEventUser.EventID)
			}
		})
	}
}

func TestInsertEventUser(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestInsertEventUser - name",
		UserName:    "TestInsertEventUser-username",
		Password:    "TestInsertEventUserpassword",
		Email:       "TestInsertEventUser@foo.bar",
		PhoneNumber: "604-555-5678",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "TestInsertEventUser - event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestInsertEventUser - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser.com",
	}

	eventUser := &models.EventUser{
		EventID: event.ID,
		UserID:  user.ID,
	}

	var tests = []struct {
		title     string
		eventUser *models.EventUser
		fail      bool
	}{
		{
			title: "fails to insert event_user with no EventID",
			eventUser: &models.EventUser{
				UserID: eventUser.UserID,
			},
			fail: true,
		},
		{
			title: "fails to insert event_user with no UserID",
			eventUser: &models.EventUser{
				EventID: eventUser.EventID,
			},
			fail: true,
		},
		{
			title:     "successfully inserts an event_user",
			eventUser: eventUser,
			fail:      false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.InsertEventUser(tt.eventUser)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDeleteEventUserByEventIDUserID(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "DeleteEventUserByEventIDUserID - name",
		UserName:    "DeleteEventUserByEventIDUserID-username",
		Password:    "DeleteEventUserByEventIDUserIDpassword",
		Email:       "DeleteEventUserByEventIDUserID@foo.bar",
		PhoneNumber: "604-955-5678",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "DeleteEventUserByEventIDUserID - event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "DeleteEventUserByEventIDUserID - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser.com",
	}

	eventUser := &models.EventUser{
		EventID: event.ID,
		UserID:  user.ID,
	}

	var tests = []struct {
		title     string
		eventUser *models.EventUser
		fail      bool
	}{
		{
			title: "fails to delete event_user with no EventID",
			eventUser: &models.EventUser{
				UserID: eventUser.UserID,
			},
			fail: true,
		},
		{
			title: "fails to delete event_user with no UserID",
			eventUser: &models.EventUser{
				EventID: eventUser.EventID,
			},
			fail: true,
		},
		{
			title: "returns success when deleting an event_user we are not a part of",
			eventUser: &models.EventUser{
				EventID: uuid.New().String(),
				UserID:  uuid.New().String(),
			},
			fail: false,
		},
		{
			title:     "successfully deletes an event_user",
			eventUser: eventUser,
			fail:      false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()
	db.Conn.Model(eventUser).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.DeleteEventUserByEventIDUserID(tt.eventUser.EventID, tt.eventUser.UserID)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				foundEventUser, _ := db.GetEventUserByEventIDUserID(tt.eventUser.EventID, tt.eventUser.UserID)
				assert.Nil(t, foundEventUser)
			}
		})
	}
}

func TestDeleteEventUserByEventID(t *testing.T) {
	userID1 := uuid.New().String()
	userID2 := uuid.New().String()
	eventID1 := uuid.New().String()
	eventID2 := uuid.New().String()

	user1 := &models.User{
		ID:          userID1,
		Name:        "DeleteEventUserByEventID - name",
		UserName:    "DeleteEventUserByEventID-username",
		Password:    "DeleteEventUserByEventIDpassword",
		Email:       "DeleteEventUserByEventID@foo.bar",
		PhoneNumber: "604-955-5678",
	}

	user2 := &models.User{
		ID:          userID2,
		Name:        "DeleteEventUserByEventID - name 2",
		UserName:    "DeleteEventUserByEventID-username 2",
		Password:    "DeleteEventUserByEventIDpassword 2",
		Email:       "DeleteEventUserByEventID@foo.bar 2",
		PhoneNumber: "604-955-5678",
	}

	event1 := &models.Event{
		ID:       eventID1,
		Name:     "DeleteEventUserByEventID - event name",
		OwnerID:  user1.ID,
		Time:     time.Now(),
		Location: "DeleteEventUserByEventID - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://DeleteEventUserByEventID.com",
	}

	event2 := &models.Event{
		ID:       eventID2,
		Name:     "DeleteEventUserByEventID - event name2",
		OwnerID:  user1.ID,
		Time:     time.Now(),
		Location: "DeleteEventUserByEventID - over there!2",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://DeleteEventUserByEventID.com",
	}

	event1User1 := &models.EventUser{
		EventID: event1.ID,
		UserID:  user1.ID,
	}

	event1User2 := &models.EventUser{
		EventID: event1.ID,
		UserID:  user2.ID,
	}

	event2User1 := &models.EventUser{
		EventID: event2.ID,
		UserID:  user1.ID,
	}

	even2User2 := &models.EventUser{
		EventID: event2.ID,
		UserID:  user2.ID,
	}

	var tests = []struct {
		title          string
		eventID        string
		eventUser1Size int
		eventUser2Size int
	}{
		{
			title:          "deletes nothing when no event exists",
			eventID:        uuid.New().String(),
			eventUser1Size: 2,
			eventUser2Size: 2,
		},
		{
			title:          "deletes only eventUsers from the passed in eventID",
			eventID:        eventID1,
			eventUser1Size: 0,
			eventUser2Size: 2,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user1).Insert()
	db.Conn.Model(user2).Insert()
	db.Conn.Model(event1).Insert()
	db.Conn.Model(event2).Insert()
	db.Conn.Model(event1User1).Insert()
	db.Conn.Model(event1User2).Insert()
	db.Conn.Model(event2User1).Insert()
	db.Conn.Model(even2User2).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.DeleteEventUserByEventID(tt.eventID)
			assert.Nil(t, err)
			eventUsers1, err := db.ListEventUsers(eventID1)
			assert.Nil(t, err)
			assert.Equal(t, tt.eventUser1Size, len(*eventUsers1))
			eventUsers2, err := db.ListEventUsers(eventID2)
			assert.Nil(t, err)
			assert.Equal(t, tt.eventUser2Size, len(*eventUsers2))
		})
	}
}

func TestListOwnedEvents(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()
	eventID2 := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestListOwnedEvents - name",
		UserName:    "TestListOwnedEvents-username",
		Password:    "TestListOwnedEventspassword",
		Email:       "TestListOwnedEvents@foo.bar",
		PhoneNumber: "604-955-5678",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "TestListOwnedEvents - event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestListOwnedEvents - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser.com",
	}

	event2 := &models.Event{
		ID:       eventID2,
		Name:     "TestListOwnedEvents - event name2",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "TestListOwnedEvents - over there!2",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser2.com",
	}

	var tests = []struct {
		title      string
		userID     string
		resultSize int
	}{
		{
			title:      "returns empty array when no owned events",
			userID:     uuid.New().String(),
			resultSize: 0,
		},
		{
			title:      "successfully returns all owned events",
			userID:     user.ID,
			resultSize: 2,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()
	db.Conn.Model(event2).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			events, err := db.ListOwnedEvents(tt.userID)
			assert.Nil(t, err)
			assert.Equal(t, tt.resultSize, len(*events))
		})
	}
}

func TestListJoinedEvents(t *testing.T) {
	userID := uuid.New().String()
	eventID := uuid.New().String()
	eventID2 := uuid.New().String()

	user := &models.User{
		ID:          userID,
		Name:        "TestListJoinedEvents - name",
		UserName:    "TestListJoinedEvents-username",
		Password:    "TestListJoinedEventspassword",
		Email:       "TestListJoinedEvents@foo.bar",
		PhoneNumber: "604-955-8678",
	}

	event := &models.Event{
		ID:       eventID,
		Name:     "DeleteEventUserByEventIDUserID - event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "DeleteEventUserByEventIDUserID - over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser.com",
	}

	event2 := &models.Event{
		ID:       eventID2,
		Name:     "DeleteEventUserByEventIDUserID - event name2",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "DeleteEventUserByEventIDUserID - over there!2",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://TestInsertEventUser2.com",
	}

	eventUser := &models.EventUser{
		EventID: event.ID,
		UserID:  user.ID,
	}

	var tests = []struct {
		title      string
		userID     string
		resultSize int
		fail       bool
	}{
		{
			title:      "successfully returns all owned events",
			userID:     user.ID,
			resultSize: 1,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user).Insert()
	db.Conn.Model(event).Insert()
	db.Conn.Model(event2).Insert()
	db.Conn.Model(eventUser).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			events, err := db.ListJoinedEvents(tt.userID)
			assert.Nil(t, err)
			assert.Equal(t, tt.resultSize, len(*events))
		})
	}
}

func TestListEventUsers(t *testing.T) {
	userID1 := uuid.New().String()
	userID2 := uuid.New().String()
	userID3 := uuid.New().String()
	userID4 := uuid.New().String()
	userID5 := uuid.New().String()
	eventID1 := uuid.New().String()
	eventID2 := uuid.New().String()
	eventID3 := uuid.New().String()

	user1 := &models.User{
		ID:          userID1,
		Name:        "TestListEventUsers - joined second name",
		UserName:    "TestListEventUsers-username",
		Password:    "TestListEventUserspassword",
		Email:       "TestListEventUsers@foo.bar",
		PhoneNumber: "604-655-8678",
	}

	user2 := &models.User{
		ID:          userID2,
		Name:        "TestListEventUsers - joined first name2",
		UserName:    "TestListEventUsers-username2",
		Password:    "TestListEventUserspassword2",
		Email:       "TestListEventUsers@foo.bar2",
		PhoneNumber: "604-655-8678",
	}

	user3 := &models.User{
		ID:          userID3,
		Name:        "TestListEventUsers - joined first name3",
		UserName:    "TestListEventUsers-username3",
		Password:    "TestListEventUserspassword3",
		Email:       "TestListEventUsers@foo.bar3",
		PhoneNumber: "604-655-8678",
	}

	user4 := &models.User{
		ID:          userID4,
		Name:        "TestListEventUsers - joined first name4",
		UserName:    "TestListEventUsers-username4",
		Password:    "TestListEventUserspassword4",
		Email:       "TestListEventUsers@foo.bar4",
		PhoneNumber: "604-655-8678",
	}

	user5 := &models.User{
		ID:          userID5,
		Name:        "TestListEventUsers - joined first name5",
		UserName:    "TestListEventUsers-username5",
		Password:    "TestListEventUserspassword5",
		Email:       "TestListEventUsers@foo.bar5",
		PhoneNumber: "604-655-8678",
	}

	eventNoMembers := &models.Event{
		ID:       eventID1,
		Name:     "TestListEventUsers - event name 1",
		OwnerID:  user1.ID,
		Time:     time.Now(),
		Location: "TestListEventUsers - over there!",
		MinUsers: 1,
		MaxUsers: 3,
		Price:    10.23,
		Link:     "http://TestListEventUsers.com",
	}

	eventWithMembersNoWaitlist := &models.Event{
		ID:       eventID2,
		Name:     "TestListEventUsers - event name 2",
		OwnerID:  user1.ID,
		Time:     time.Now(),
		Location: "TestListEventUsers - over there!",
		MinUsers: 1,
		MaxUsers: 3,
		Price:    10.23,
		Link:     "http://TestListEventUsers.com",
	}

	eventWithMembersAndWaitlist := &models.Event{
		ID:       eventID3,
		Name:     "TestListEventUsers - event name 3",
		OwnerID:  user1.ID,
		Time:     time.Now(),
		Location: "TestListEventUsers - over there!",
		MinUsers: 1,
		MaxUsers: 3,
		Price:    10.23,
		Link:     "http://TestListEventUsers.com",
	}

	curTime := time.Now().UTC()
	oneHours := time.Now().UTC().Add(time.Hour * time.Duration(1))
	twoHours := time.Now().UTC().Add(time.Hour * time.Duration(2))
	threeHours := time.Now().UTC().Add(time.Hour * time.Duration(3))
	fourHours := time.Now().UTC().Add(time.Hour * time.Duration(4))
	fiveHours := time.Now().UTC().Add(time.Hour * time.Duration(5))
	eventUser1 := &models.EventUser{
		EventID:   eventWithMembersNoWaitlist.ID,
		UserID:    user1.ID,
		CreatedAt: fiveHours, //joining 5 hours later
	}
	eventUser2 := &models.EventUser{
		EventID:   eventWithMembersNoWaitlist.ID,
		UserID:    user2.ID,
		CreatedAt: curTime,
	}

	eventUserWaitlist1 := &models.EventUser{
		EventID:   eventWithMembersAndWaitlist.ID,
		UserID:    user1.ID,
		CreatedAt: threeHours,
	}

	eventUserWaitlist2 := &models.EventUser{
		EventID:   eventWithMembersAndWaitlist.ID,
		UserID:    user2.ID,
		CreatedAt: twoHours,
	}

	eventUserWaitlist3 := &models.EventUser{
		EventID:   eventWithMembersAndWaitlist.ID,
		UserID:    user3.ID,
		CreatedAt: fourHours, //joining 5 hours later
	}

	eventUserWaitlist4 := &models.EventUser{
		EventID:   eventWithMembersAndWaitlist.ID,
		UserID:    user4.ID,
		CreatedAt: oneHours, //joining 5 hours later
	}

	eventUserWaitlist5 := &models.EventUser{
		EventID:   eventWithMembersAndWaitlist.ID,
		UserID:    user5.ID,
		CreatedAt: curTime, //joining 5 hours later
	}

	var tests = []struct {
		title      string
		eventID    string
		resultSize int
		inSize     int
		fail       bool
	}{
		{
			title:      "returns error when event doesnt exist",
			eventID:    uuid.New().String(),
			resultSize: 0,
			inSize:     0,
			fail:       true,
		},
		{
			title:      "returns empty list when event has no members",
			eventID:    eventNoMembers.ID,
			resultSize: 0,
			inSize:     0,
			fail:       false,
		},
		{
			title:      "returns list of users when event has members, none in wait list",
			eventID:    eventWithMembersNoWaitlist.ID,
			resultSize: 2,
			inSize:     2,
			fail:       false,
		},
		{
			title:      "returns list of users when event has members in and out of wait list",
			eventID:    eventWithMembersAndWaitlist.ID,
			resultSize: 5,
			inSize:     int(eventWithMembersAndWaitlist.MaxUsers),
			fail:       false,
		},
	}

	// setup db
	db, _ := NewDB()
	db.Conn.Model(user1).Insert()
	db.Conn.Model(user2).Insert()
	db.Conn.Model(user3).Insert()
	db.Conn.Model(user4).Insert()
	db.Conn.Model(user5).Insert()
	db.Conn.Model(eventNoMembers).Insert()
	db.Conn.Model(eventWithMembersNoWaitlist).Insert()
	db.Conn.Model(eventWithMembersAndWaitlist).Insert()
	db.Conn.Model(eventUser1).Insert()
	db.Conn.Model(eventUser2).Insert()
	_, err := db.Conn.Model(eventUserWaitlist1).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Conn.Model(eventUserWaitlist2).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Conn.Model(eventUserWaitlist3).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Conn.Model(eventUserWaitlist4).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Conn.Model(eventUserWaitlist5).Insert()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			eventUsers, err := db.ListEventUsers(tt.eventID)
			if tt.fail {
				assert.NotNil(t, err)
				assert.Nil(t, eventUsers)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.resultSize, len(*eventUsers))
				for i, eventUser := range *eventUsers {
					if i < tt.inSize {
						assert.True(t, eventUser.IsIn)
					} else {
						assert.False(t, eventUser.IsIn)
					}
				}
			}
		})
	}
}
