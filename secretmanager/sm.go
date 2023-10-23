package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/f3rcho/twitterGo/awsgo"
	"github.com/f3rcho/twitterGo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secretValues models.Secret
	fmt.Println("Getting secret..." + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Conf)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretValues, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretValues)
	fmt.Println("Reading Secret OK" + secretName)
	return secretValues, nil
}
