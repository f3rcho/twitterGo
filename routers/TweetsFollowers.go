package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/models"
)

func GetTweetsFollowers(request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400
	IDUser := claim.ID.Hex()

	page := request.QueryStringParameters["page"]
	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Page must be greater than 0"
		return r
	}
	tweets, OK := db.GetTweetsFollowers(IDUser, pag)
	if !OK {
		r.Message = "Error getting tweets"
		return r
	}

	resJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error formating tweets"
		return r
	}

	r.Status = 200
	r.Message = string(resJson)
	return r
}
