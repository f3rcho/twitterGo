package routers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func Relation(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "ID param required"
		return r
	}

	var re models.Relation

	re.UsuerId = claim.ID.Hex()
	re.UserRelationId = ID

	status, err := db.InsertRelation(re)
	if err != nil {
		r.Message = "Error trying to insert relation " + err.Error()
		return r
	}
	if !status {
		r.Message = "Relation not inserted"
		return r
	}
	r.Status = 200
	r.Message = "Relation inserted"
	return r
}
