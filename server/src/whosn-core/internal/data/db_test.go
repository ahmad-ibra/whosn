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
