package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"tasks/models"
	"tasks/tools"

	_ "github.com/go-sql-driver/mysql"
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

func UpdateTask(t models.Task) error {
	fmt.Println("Start UPDATE")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "UPDATE tasks SET "

	sentence = tools.ArmoSentencia(sentence, "task_title", "S", 0, 0, t.TaskTitle)
	sentence = tools.ArmoSentencia(sentence, "task_description", "S", 0, 0, t.TaskDescription)
	sentence = tools.ArmoSentencia(sentence, "is_done", "N", t.TaskDone, 0, "")

	sentence += " WHERE Task_Id = "+strconv.Itoa(t.TaskID)

	fmt.Println(sentence)

	_, err = Db.Exec(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Task > Successfull")

	return nil
}