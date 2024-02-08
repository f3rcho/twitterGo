package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RetriveTweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"userId" json:"userid,omitempty"`
	UserRelationId string             `bson:"userrelationid" json:"userrelationid"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
