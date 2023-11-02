package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(r models.Relation) bool {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("relations")

	condition := bson.M{
		"userid":         r.UsuerId,
		"userrelationid": r.UserRelationId,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	return err == nil
}
