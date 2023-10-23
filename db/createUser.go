package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
)

func CreateUser(u models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoClient.Database(DatabaseName)
	colection := db.Collection("users")

	u.Password, _ = Encrypt(u.Password)

	ressult, err := colection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjectID, _ := ressult.InsertedID.(string)
	return ObjectID, true, nil
}
