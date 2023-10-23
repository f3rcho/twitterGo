package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func IsUserExists(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoClient.Database(DatabaseName)
	colection := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := colection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
