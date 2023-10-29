package routers

import (
	"context"
	"encoding/json"

	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func UpdateUser(ctx context.Context, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	var u models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		r.Message = "Invalid Data" + err.Error()
	}
	var status bool

	status, err := db.UpdateUser(u.claim.ID.Hex())
	if err != nil {
		r.Message = "Error updating user" + err.Error()
		return r
	}
	if !status {
		r.Message = "Can not update user"
		return r
	}
	r.Status = 200
	r.Message = "User updated"
	return r
}
