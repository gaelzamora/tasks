package secrets

import (
	"encoding/json"
	"fmt"
	"tasks/awsgo"
	"tasks/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(" > Get Secret "+nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Println(" > Read Secret OK "+nameSecret)

	return dataSecret, nil
	
}