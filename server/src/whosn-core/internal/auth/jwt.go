package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// TODO: store this in a .env
var jwtSigningKey = []byte("supersecretkey")

type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtClaim := JWTClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "whosn-core",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString(jwtSigningKey)
}

func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})

	var userID string
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		userID = claims.UserID
	} else {
		ll := log.WithFields(log.Fields{"function": "ValidateToken", "error": err})
		ll.Warnf("Invalid token")
		return "", fmt.Errorf("token invalid. %v", err)
	}

	return userID, nil
}
