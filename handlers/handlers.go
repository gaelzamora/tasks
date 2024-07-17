package handlers

import (
	"fmt"
	"strconv"
	"tasks/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("I'm going to proccess: "+path+" > "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	if len(path) < 4 {
		fmt.Println(path)
		return 400, "Path too short"
	}

	switch path[0:4] {
	case "task":
		return TaskProcess(body, path, method, idn, request)
	}

	return 400, "Method invalid"
}

func TaskProcess(body string, path string, method string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
		switch method {
		case "POST":
			return routers.InsertTask(body)
		case "PUT":
			return routers.UpdateProduct(body, id)
		}

	return 400, "Method invalid"
}