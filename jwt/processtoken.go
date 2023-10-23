package jwt

import (
	"errors"
	"strings"

	"github.com/f3rcho/twitterGo/models"
	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myPass := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("invalid format token")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})
	if err == nil {
		// check with DB
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invalid token")
	}
	return &claims, false, string(""), nil
}
