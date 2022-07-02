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

	var tests = []struct {
		title string
		user  *models.User
		fail  bool
	}{
		{
			title: "successfully creates a user",
			user: &models.User{
				Name:        "some name",
				UserName:    "someUserName",
				Password:    "password",
				Email:       "email@foo.bar",
				PhoneNumber: "604-555-5555",
			},
			fail: false,
		},
	}

	db, err := NewDB()
	assert.Nil(t, err)

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
