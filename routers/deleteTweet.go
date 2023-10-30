package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/models"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID is required"
		return r
	}
	err := db.DeleteOne(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "An error trying to delete the tweet" + err.Error()
		return r
	}

	r.Message = "Tweet deleted successfully"
	r.Status = 200
	return r
}
