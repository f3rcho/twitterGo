# twitter app
  Api created with go
# Go
## mongoAtlas
  For DB
## aws cloud
  - Run lambdas
  - upload the app image
  - s3 (twitterg) to save the media
  
### S3 policy
```json
{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Effect":"Allow",
         "Principal": "*",
         "Action":[
            "s3:GetObject",
            "s3:ListBucket"
         ],
         "Resource": [
            "arn:aws:s3:::twitterg",
            "arn:aws:s3:::twitterg/*"
        ]
      }
   ]
}
```

## Secret Manager
```go
// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:   
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	secretName := "twitter"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	// Your code goes here.
} 
```

## Api gateway 
- create api gateway
- create resource, enable proxy and cors. resourceRoute {twitter+}
- In actions, implement API. Name the step. 
- In config, binary media types, add multipart/form-data and image/jpg