package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/jwt"
	"github.com/f3rcho/twitterGo/models"
)

func Login(ctx context.Context) models.ResposenAPI {
	var u models.User
	var r models.ResposenAPI
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "Invalid User or Password" + err.Error()
		return r
	}
	if len(u.Email) == 0 {
		r.Message = ""
		return r
	}
	userData, exists := db.TryLogin(u.Email, u.Password)
	if !exists {
		r.Message = "Invalid User or Password"
		return r
	}
	jwtKey, err := jwt.GenerateJWT(ctx, userData)
	if err != nil {
		r.Message = "An error has occured trying to generate token" + err.Error()
		return r
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}
	token, errToken := json.Marshal(resp)
	if errToken != nil {
		r.Message = "An error has occured trying to generate token" + errToken.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}
	r.Status = 200
	r.Message = string(token)
	r.CustomResponse = res

	return r
}
