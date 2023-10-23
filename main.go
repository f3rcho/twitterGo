package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/f3rcho/twitterGo/awsgo"
	"github.com/f3rcho/twitterGo/secretmanager"
)

func main() {
	lambda.Start(ExecuteLambda)

}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()

	if !validateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Missing parameters SecretName, BucketName, UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error reading secret" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	return res, nil
}

func validateParams() bool {
	_, bringParams := os.LookupEnv("SecretName")
	if !bringParams {
		return bringParams
	}
	_, bringParams = os.LookupEnv("BucketName")
	if !bringParams {
		return bringParams
	}
	_, bringParams = os.LookupEnv("UrlPrefix")
	if !bringParams {
		return bringParams
	}
	return bringParams
}
