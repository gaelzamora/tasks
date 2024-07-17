package routers

import (
	"encoding/json"
	"strconv"
	"tasks/database"
	"tasks/models"
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