package models

type Relation struct {
	UsuerId        string `bson:"userId" json:"userid"`
	UserRelationId string `bson:"userrelationid" json:"userrelationid"`
}
