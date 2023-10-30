package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTweet, bool) {
	ctx := context.TODO()

	db := MongoClient.Database(DatabaseName)
	col := db.Collection("tweets")

	var result []*models.GetTweet
	condition := bson.M{
		"userId": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		return result, false
	}
	for cursor.Next(ctx) {
		var tweet models.GetTweet
		err := cursor.Decode(&tweet)
		if err != nil {
			return result, false
		}
		result = append(result, &tweet)
	}
	return result, true
}
