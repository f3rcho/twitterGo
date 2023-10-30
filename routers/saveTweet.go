package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/f3rcho/twitterGo/models"
)

func SaveTweet(ctx context.Context, claim models.Claim) models.ResposenAPI {
	var message models.Tweet
	var r models.ResposenAPI
	r.Status = 400
	IDUser := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		r.Message = "Error trying to decode body" + err.Error()
		return r
	}
	tweet := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := db.createTweet(tweet)
	if err != nil {
		r.Message = "Error trying to create tweet" + err.Error()
		return r
	}
	if !status {
		r.Message = "Tweet could not be created"
		return r
	}
	r.Status = 200
	r.Message = "Tweet created successfully"
	return r

}
