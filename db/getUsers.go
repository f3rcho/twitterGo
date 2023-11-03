package db

import (
	"context"
	"fmt"

	"github.com/f3rcho/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(ID string, page int64, search string, typeUser string) ([]*models.User, bool) {
	ctx := context.TODO()
	db := MongoClient.Database(DatabaseName)
	col := db.Collection("users")

	var result []*models.User
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip((page - 1) * 20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cursor, err := col.Find(ctx, query, options)
	if err != nil {
		return result, false
	}

	var include bool

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Decode error: " + err.Error())
			return result, false
		}

		var re models.Relation
		re.UsuerId = ID
		re.UserRelationId = user.ID.Hex()

		include = false

		found := GetRelation(re)
		if typeUser == "new" && !found {
			include = true
		}
		if typeUser == "follow" && found {
			include = true
		}

		if re.UserRelationId == ID {
			include = false
		}
		if include {
			user.Password = ""
			result = append(result, &user)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println("Cursor error: " + err.Error())
		return result, false
	}
	cursor.Close(ctx)
	return result, true
}
