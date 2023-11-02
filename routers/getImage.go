package routers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/f3rcho/twitterGo/awsgo"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

func GetImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "The ID is required"
		return r
	}

	user, err := db.FindOne(ID)
	if err != nil {
		r.Message = "User not found " + err.Error()
		return r
	}

	var filename string
	switch uploadType {
	case "A":
		filename = user.Avatar
	case "B":
		filename = user.Banner
	}
	fmt.Println("Filename " + filename)
	svc := s3.NewFromConfig(awsgo.Conf)

	file, err := downloadFromS3(ctx, svc, filename)

	if err != nil {
		r.Status = 500
		r.Message = "Error downloading file from S3 " + err.Error()
		return r
	}

	r.CustomResponse = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       file.String(),
		Headers: map[string]string{
			"Content-Type":        "application/octet-stream",
			"Content-Disposition": "attachment; filename=" + filename,
		},
	}
	return r
}

func downloadFromS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	defer obj.Body.Close()
	fmt.Println("BucketName = " + bucket)

	file, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(file)

	return buffer, nil
}
