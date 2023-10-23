package db

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
