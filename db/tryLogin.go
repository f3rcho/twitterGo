package db

import (
	"github.com/f3rcho/twitterGo/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := IsUserExists(email)
	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
