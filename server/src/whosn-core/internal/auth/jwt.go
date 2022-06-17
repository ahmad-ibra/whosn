package auth

import (
	"fmt"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID string, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtClaim := JWTClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    config.SvcName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString([]byte(jwtKey))
}

func ValidateToken(signedToken string, jwtKey string) (string, error) {
	jwtSigningKey := []byte(jwtKey)

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
