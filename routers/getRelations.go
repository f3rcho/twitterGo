package routers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func GetRelations(request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID is required"
		return r
	}
	var re models.Relation
	re.UsuerId = claim.ID.Hex()
	re.UserRelationId = ID

	var res models.ResponseRelation

	isRelation := db.GetRelation(re)
	if !isRelation {
		res.Status = false
	} else {
		res.Status = true
	}

	resJson, err := json.Marshal(isRelation)
	if err != nil {
		r.Status = 500
		r.Message = "Error formationg response: " + err.Error()
		return r
	}
	r.Status = 200
	r.Message = string(resJson)
	return r
}
