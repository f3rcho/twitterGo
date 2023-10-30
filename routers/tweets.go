package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func GetTweets(request events.APIGatewayProxyRequest) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["page"]

	if len(ID) < 1 {
		r.Message = "the param ID is required"
		return r
	}
	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Page must be greater than 0"
	}
	tweets, OK := db.GetTweets(ID, int64(pag))
	if !OK {
		r.Message = "Error reading tweets"
		return r
	}
	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error parsing tweets"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
