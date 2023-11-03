package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func GetUsers(request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IdUser := claim.ID.Hex()

	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Page must be greater than 0"
	}

	users, OK := db.GetUsers(IdUser, int64(pag), search, typeUser)

	if !OK {
		r.Message = "Error getting users"
		return r
	}
	respJson, err := json.Marshal(users)
	if err != nil {
		r.Status = 500
		r.Message = "Error parsing users"
		return r
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
