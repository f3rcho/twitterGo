package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/f3rcho/twitterGo/awsgo"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/handlers"
	"github.com/f3rcho/twitterGo/models"
	"github.com/f3rcho/twitterGo/secretmanager"
)

func main() {
	lambda.Start(ExecuteLambda)

}

const KEY_HEADER = "Content-Type"
const VALUE_HEADER = "application/json"

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()

	if !validateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Missing parameters SecretName, BucketName, UrlPrefix",
			Headers: map[string]string{
				KEY_HEADER: VALUE_HEADER,
			},
		}
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error reading secret" + err.Error(),
			Headers: map[string]string{
				KEY_HEADER: VALUE_HEADER,
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitter"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	// check db connection
	err = db.ConnectDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connecting to database" + err.Error(),
			Headers: map[string]string{
				KEY_HEADER: VALUE_HEADER,
			},
		}
		return res, nil
	}

	responseAPI := handlers.Handlers(awsgo.Ctx, request)
	if responseAPI.CustomResponse == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: responseAPI.Status,
			Body:       responseAPI.Message,
			Headers: map[string]string{
				KEY_HEADER: VALUE_HEADER,
			},
		}
		return res, nil
	} else {
		return responseAPI.CustomResponse, nil
	}
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
