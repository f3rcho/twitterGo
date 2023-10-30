package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/jwt"
	"github.com/f3rcho/twitterGo/models"
	"github.com/f3rcho/twitterGo/routers"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResposenAPI {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string))

	var r models.ResposenAPI
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
		case "login":
			return routers.Login(ctx)
		case "tweet":
			return routers.SaveTweet(ctx, claim)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return routers.GetProfile(request)
		case "tweets":
			return routers.GetTweets(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return routers.UpdateUser(ctx, claim)
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "tweet":
			return routers.DeleteTweet(ctx, claim)
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
	claim, Ok, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtsign")).(string))
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
