package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/models"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResposeAPI {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string))

	var r models.ResposeAPI
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

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
