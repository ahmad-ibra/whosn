package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {

	var tests = []struct {
		title  string
		userID string
	}{
		{
			title:  "check jwt token can be successfully generated and validated",
			userID: "5121efb2-c49a-405d-b804-05bd04a6a601",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			tokenString, err := GenerateJWT(tt.userID)
			assert.Nil(t, err)

			userID, err := ValidateToken(tokenString)
			assert.Nil(t, err)
			assert.Equal(t, tt.userID, userID)
		})
	}
}
