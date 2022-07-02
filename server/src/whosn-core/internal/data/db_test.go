package data

import (
	"os"
	"testing"

	"github.com/Ahmad-Ibra/whosn-core/internal/data/models"
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

	// clean table
	db.Conn.Model(user).Where("user_name = ?", user.UserName).Delete()
	db.Conn.Model(user).Where("user_name = ?", duplicateUser.UserName).Delete()

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

	// clean table
	db.Conn.Model(user).Where("user_name = ?", user.UserName).Delete()

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
