package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"tasks/database"
	"tasks/handlers"
	"tasks/awsgo"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Loading Lambda")
	awsgo.StartAWS()

	if !ParametersValid() {
		panic("Error en los parametros, debe enviar 'SecretName', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	database.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string {
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}

	return res, nil
}

func ParametersValid() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}

	fmt.Println(getParam)

	_, getParam = os.LookupEnv("UrlPrefix")

	if !getParam {
		return getParam
	}

	return getParam
}
