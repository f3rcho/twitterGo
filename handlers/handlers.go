package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/jwt"
	"github.com/f3rcho/twitterGo/models"
	"github.com/f3rcho/twitterGo/routers"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResposeAPI {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string))

	var r models.ResposeAPI
	r.Status = 400

	isOk, statusCode, msg, claim := validateAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Invalid Method"
	return r
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token required", models.Claim{}
	}
	claim, Ok, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("sign")).(string))
	if !Ok {
		if err != nil {
			fmt.Println("Token Error " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Token Error " + msg)
			return false, 401, msg, models.Claim{}
		}
	}
	fmt.Println("Token OK")
	return true, 200, "", *claim
}
