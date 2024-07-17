package database

import (
	"database/sql"
	"fmt"
	"tasks/models"
	"tasks/tools"
)

func InsertTask(t models.Task) (int64, error) {
	fmt.Println("Start Insert")

	err := DbConnect()

	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentence := "INSERT INTO tasks (task_title"

	if len(t.TaskDescription) > 0 {
		sentence += ", task_description"
	}

	sentence += ") VALUES ('"+tools.EscapeString(t.TaskTitle)+"'"

	if len(t.TaskDescription) > 0 {
		sentence += ",'"+tools.EscapeString(t.TaskDescription)+"'"
	}

	sentence += ")"

	var result sql.Result

	result, err = Db.Exec(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Task > Succesfull")
	return LastInsertId, nil


}