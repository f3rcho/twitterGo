package db

import (
	"context"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.User, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("users")

	updatedUser := make(map[string]interface{})
	if len(u.Name) > 0 {
		updatedUser["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		updatedUser["lastName"] = u.LastName
	}
	if len(u.Email) > 0 {
		updatedUser["birthDay"] = u.BirthDay
	}
	if len(u.Avatar) > 0 {
		updatedUser["avatar"] = u.Avatar
	}
	if len(u.Location) > 0 {
		updatedUser["location"] = u.Location
	}
	if len(u.Biography) > 0 {
		updatedUser["biography"] = u.Biography
	}
	if len(u.Website) > 0 {
		updatedUser["website"] = u.Website
	}
	if len(u.Banner) > 0 {
		updatedUser["banner"] = u.Banner
	}

	updateString := bson.M{
		"$set": updatedUser,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
