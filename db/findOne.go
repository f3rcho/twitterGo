package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOne(ID string) (models.User, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}
	return profile, nil

}
