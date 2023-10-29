package models

import "github.com/aws/aws-lambda-go/events"

type ResposenAPI struct {
	Status         int
	Message        string
	CustomResponse *events.APIGatewayProxyResponse
}
