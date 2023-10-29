package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func GetProfile(request events.APIGatewayProxyRequest) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	fmt.Println("Getting profile...")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID is required"
		return r
	}

	profile, err := db.FindOne(ID)
	if err != nil {
		r.Message = "An error has occured getting the profile" + err.Error()
		return r
	}
	respJson, err := json.Marshal(profile)
	if err != nil {
		r.Status = 500
		r.Message = "Error formating user data as JSON" + err.Error()
		return r
	}
	r.Status = 200
	r.Message = string(respJson)
	return r
}
