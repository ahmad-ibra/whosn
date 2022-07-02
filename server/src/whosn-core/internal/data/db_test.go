package data

import (
	"os"
	"testing"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("ENV", "test")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "dev")
	os.Setenv("POSTGRES_PASSWORD", "pass")
	os.Setenv("POSTGRES_DBNAME", "whosn")
}

func cleanTables(db *pg.DB) {
	var users []models.User
	db.Model(&users).Select()
	db.Model(&users).Delete()

	var events []models.Event
	db.Model(&events).Select()
	db.Model(&events).Delete()

	var eventUsers []models.EventUser
	db.Model(&eventUsers).Select()
	db.Model(&eventUsers).Delete()
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
		Name:        "some name",
		UserName:    "someUserName",
		Password:    "password",
		Email:       "email@foo.bar",
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
				UserName:    "someUserName",
				Password:    "password",
				Email:       "email@foo.bar",
				PhoneNumber: "604-555-5555",
			},
			fail: true,
		},
		{
			title: "fails to insert user with no UserName",
			user: &models.User{
				Name:        "some name",
				Password:    "password",
				Email:       "email@foo.bar",
				PhoneNumber: "604-555-5555",
			},
			fail: true,
		},
		{
			title: "fails to insert user with no Password",
			user: &models.User{
				Name:        "some name",
				UserName:    "someUserName",
				Email:       "email@foo.bar",
				PhoneNumber: "604-555-5555",
			},
			fail: true,
		},
		{
			title: "fails to insert user with no Email",
			user: &models.User{
				Name:        "some name",
				UserName:    "someUserName",
				Password:    "password",
				PhoneNumber: "604-555-5555",
			},
			fail: true,
		},
		{
			title: "fails to insert user with no PhoneNumber",
			user: &models.User{
				Name:     "some name",
				UserName: "someUserName",
				Password: "password",
				Email:    "email@foo.bar",
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
	cleanTables(db.Conn)

	// insert duplicate
	db.Conn.Model(duplicateUser).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.InsertUser(tt.user)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				db.Conn.Model(tt.user).Where("id = ?", tt.user.ID).Delete()
			}
		})
	}
}

func TestGetUserByUserName(t *testing.T) {
	user := &models.User{
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
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
	cleanTables(db.Conn)

	// insert user
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
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
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
			id:    "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
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
			id:    "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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

	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "username",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
		CreatedAt:   createTime,
		UpdatedAt:   createTime,
	}

	updatedUser := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "new name",
		UserName:    "updatedUsername",
		Password:    "newPassword",
		Email:       "newEmail@foo.bar",
		PhoneNumber: "604-555-9999",
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
			id:      "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "username",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	event := &models.Event{
		Name:     "event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://somelink.com",
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
				Time:     time.Now(),
				Location: "over there!",
				MinUsers: 1,
				MaxUsers: 4,
				Price:    10.23,
				Link:     "http://somelink.com",
			},
			fail: true,
		},
		{
			title: "fails to insert event with no OwnerID",
			event: &models.Event{
				Name:     "event name",
				Time:     time.Now(),
				Location: "over there!",
				MinUsers: 1,
				MaxUsers: 4,
				Price:    10.23,
				Link:     "http://somelink.com",
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Time",
			event: &models.Event{
				Name:     "event name",
				OwnerID:  user.ID,
				Location: "over there!",
				MinUsers: 1,
				MaxUsers: 4,
				Price:    10.23,
				Link:     "http://somelink.com",
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Location",
			event: &models.Event{
				Name:     "event name",
				OwnerID:  user.ID,
				Time:     time.Now(),
				MinUsers: 1,
				MaxUsers: 4,
				Price:    10.23,
				Link:     "http://somelink.com",
			},
			fail: true,
		},
		{
			title: "fails to insert event with MinUsers > MaxUsers",
			event: &models.Event{
				Name:     "event name",
				OwnerID:  user.ID,
				Time:     time.Now(),
				Location: "over there!",
				MinUsers: 6,
				MaxUsers: 2,
				Price:    10.23,
				Link:     "http://somelink.com",
			},
			fail: true,
		},
		{
			title: "fails to insert event with no Link",
			event: &models.Event{
				Name:     "event name",
				OwnerID:  user.ID,
				Time:     time.Now(),
				Location: "over there!",
				MinUsers: 3,
				MaxUsers: 7,
				Price:    10.23,
			},
			fail: true,
		},
		{
			title: "successfully inserts event with no Price",
			event: &models.Event{
				Name:     "event name",
				OwnerID:  user.ID,
				Time:     time.Now(),
				Location: "over there!",
				MinUsers: 3,
				MaxUsers: 7,
				Link:     "http://somelink.com",
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
	cleanTables(db.Conn)

	// insert duplicate
	db.Conn.Model(user).Insert()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := db.InsertEvent(tt.event)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				db.Conn.Model(tt.event).Where("id = ?", tt.event.ID).Delete()
			}
		})
	}
}

func TestGetEventByID(t *testing.T) {
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	event := &models.Event{
		ID:       "9b73daa3-c8e5-4a94-b638-4877f5edcc4f",
		Name:     "event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://somelink.com",
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
			id:    "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	event := &models.Event{
		ID:       "9b73daa3-c8e5-4a94-b638-4877f5edcc4f",
		Name:     "event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://somelink.com",
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
			id:    "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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

	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "username",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	event := &models.Event{
		ID:        "9b73daa3-c8e5-4a94-b638-4877f5edcc4f",
		Name:      "event name",
		OwnerID:   user.ID,
		Time:      createTime,
		Location:  "over there!",
		MinUsers:  1,
		MaxUsers:  4,
		Price:     10.23,
		Link:      "http://somelink.com",
		CreatedAt: createTime,
		UpdatedAt: createTime,
	}

	updatedEvent := &models.Event{
		ID:        "9b73daa3-c8e5-4a94-b638-4877f5edcc4f",
		Name:      "event name",
		OwnerID:   user.ID,
		Time:      createTime,
		Location:  "over there!",
		MinUsers:  1,
		MaxUsers:  4,
		Price:     10.23,
		Link:      "http://somelink.com",
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
			id:       "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)

	// insert user
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
	user := &models.User{
		ID:          "f1777653-0378-4b75-b8a2-4305b170917d",
		Name:        "some name",
		UserName:    "testGetUserByUserName",
		Password:    "password",
		Email:       "email@foo.bar",
		PhoneNumber: "604-555-5555",
	}

	event := &models.Event{
		ID:       "9b73daa3-c8e5-4a94-b638-4877f5edcc4f",
		Name:     "event name",
		OwnerID:  user.ID,
		Time:     time.Now(),
		Location: "over there!",
		MinUsers: 1,
		MaxUsers: 4,
		Price:    10.23,
		Link:     "http://somelink.com",
	}

	eventUser := &models.EventUser{
		ID:      "8693b1eb-96b9-4cc7-bd22-55130a98588d",
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
			eventID: "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
			userID:  user.ID,
			fail:    true,
		},
		{
			title:   "fails to find event_user if userID is not in db",
			eventID: event.ID,
			userID:  "8d5db8fa-85bb-44e1-9a93-4fdd3c866ccc",
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
	cleanTables(db.Conn)
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
