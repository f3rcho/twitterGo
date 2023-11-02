package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
)

func InsertRelation(r models.Relation) (bool, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relations")
	_, err := col.InsertOne(ctx, r)
	if err != nil {
		return false, err
	}
	return true, nil
}
