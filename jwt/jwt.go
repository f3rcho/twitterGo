package jwt

import (
	"context"
	"time"

	"github.com/f3rcho/twitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, u models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)
	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastName":  u.LastName,
		"birthDay":  u.BirthDay,
		"location":  u.Location,
		"webSite":   u.Website,
		"biography": u.Biography,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
