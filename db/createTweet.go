package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTweet(t models.SaveTweet) (string, bool, error) {
	ctx := context.TODO()

	db := MongoClient.Database(DatabaseName)
	col := db.Collection("tweets")

	tweet := bson.M{
		"userId":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, tweet)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
