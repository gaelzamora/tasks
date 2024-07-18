package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tasks/database"
	"tasks/models"

	"github.com/aws/aws-lambda-go/events"
)

func InsertTask(body string) (int, string) {
	var t models.Task

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Data error"+err.Error()
	}

	if len(t.TaskTitle)==0 {
		return 400, "You have specific Task's name"
	}

	result, err2 := database.InsertTask(t)

	if err2 != nil {
		return 400, "Error to attemp register a task"+t.TaskTitle+err2.Error()
	}

	return 200, "{ TaskID: "+strconv.Itoa(int(result))+" }"
}

func UpdateProduct(body string, id int) (int, string) {
	var t models.Task

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error in data "+err.Error()
	}

	t.TaskID = id
	err2 := database.UpdateTask(t)

	if err2 != nil {
		return 400, "There's an error to doing UPDATE of Task"+strconv.Itoa(id)+" > "+err2.Error()
	}

	return 200, "Update OK"
}

func SelectTasks(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Task
	var page, pageSize int
	var orderType, orderField string
	fmt.Println("Enter to Select Task")

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])
	pageSize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"]
	orderField = param["orderField"]

	if !strings.Contains("ITDFPCS", orderField) {
		orderField=""
	}

	var choice string
	if len(param["taskID"]) > 0 {
		choice="P"
		t.TaskID, _ = strconv.Atoi(param["taskID"])
	}

	result, err2 := database.SelectTask(t, choice, page, pageSize, orderType, orderField)

	fmt.Println(result)

	if err2 != nil {
		return 400, "Error to capture attemp results " + choice + err2.Error()
	}

	Task, err3 := json.Marshal(result)

	fmt.Println(Task)

	if err3 != nil {
		return 400, "Error to convert attemp JSON to Tasks"
	}

	return 200, string(Task)
}

func DeleteTask(id int) (int, string) {
	err := database.DeleteTask(id)

	if err != nil {
		return 400, "Error to Delete task"+err.Error()
	}

	return 200, "DeleteTask OK"
}