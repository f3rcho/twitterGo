package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Conf aws.Config
var err error

func InitAWS() {
	Ctx = context.TODO()
	Conf, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("Error loading AWS config .aws/config" + err.Error())
	}
}
