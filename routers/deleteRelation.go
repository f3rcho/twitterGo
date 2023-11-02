package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/f3rcho/twitterGo/models"
)

func DeleteRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
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

	status, err := re.DeleteRelation(re)
	if err != nil {
		r.Message = "Error trying to delete relation " + err.Error()
		return r
	}
	if !status {
		r.Message = "Relation not deleted"
		return r
	}
	r.Status = 200
	r.Message = "Relation deleted"
	return r

}
