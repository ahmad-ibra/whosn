package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {

	var tests = []struct {
		title      string
		user       *User
		hashPword  string
		checkPword string
		fail       bool
	}{
		{
			title:      "check password returns error when incorrect password is provided",
			user:       &User{},
			hashPword:  "123456789",
			checkPword: "abcdefg",
			fail:       true,
		},
		{
			title:      "check password succeeds when correct password is provided",
			user:       &User{},
			hashPword:  "123456789",
			checkPword: "123456789",
			fail:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			err := tt.user.HashPassword(tt.hashPword)
			assert.Nil(t, err)

			err = tt.user.CheckPassword(tt.checkPword)
			if tt.fail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
