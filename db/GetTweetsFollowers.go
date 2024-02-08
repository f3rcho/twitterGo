package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTweetsFollowers(ID string, page int) ([]models.RetriveTweetsFollowers, bool) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{
		"userId": ID,
	}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localfield":   "userrelationid",
			"foreignfield": "userId",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	var results []models.RetriveTweetsFollowers
	cursor, err := col.Aggregate(ctx, conditions)
	if err != nil {
		return results, false
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}

	return results, true

}
