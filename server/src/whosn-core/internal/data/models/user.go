package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// User holds data coming from the users table
type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateUserBody holds the body of the request to create a user
type CreateUserBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) ConstructUpdate(original *User) error {
	user.ID = original.ID
	if user.Name == "" {
		user.Name = original.Name
	}
	if user.UserName == "" {
		user.UserName = original.UserName
	}
	if user.Password == "" {
		user.Password = original.Password
	} else {
		if err := user.HashPassword(user.Password); err != nil {
			log.Warn("Failed to hash password")
			return err
		}
	}
	if user.Email == "" {
		user.Email = original.Email
	}
	if user.PhoneNumber == "" {
		user.PhoneNumber = original.PhoneNumber
	}
	user.CreatedAt = original.CreatedAt
	user.UpdatedAt = time.Now().UTC()

	return nil
}
