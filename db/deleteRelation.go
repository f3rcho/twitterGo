package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
)

func DeleteRelation(r models.Relation) (bool, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, r)
	if err != nil {
		return false, err
	}
	return true, nil
}
