package database

import (
	"database/sql"
	"fmt"
	"log"
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

func SelectTask(t models.Task, choice string, page int, pageSize int, orderType string, orderField string) (models.TaskResp, error) {
	fmt.Println("Start SelectTask")
	var Resp models.TaskResp
	var task []models.Task

	err := DbConnect()
	if err != nil {
		return Resp, err
	}
	defer Db.Close()

	log.Println("ID: "+strconv.Itoa(t.TaskID))
	log.Println("Choice: "+choice)

	var sentence string
	var sentenceCount string
	var where, limit string

	sentence = "SELECT Task_Id, task_title, task_description, is_done FROM tasks "
	sentenceCount = "SELECT count(*) as registers FROM tasks "

	switch choice {
	case "P":
		where = " WHERE Task_Id = "+strconv.Itoa(t.TaskID)
	}

	sentenceCount += where
	log.Println("Sentence Couint: "+sentenceCount)

	var rows *sql.Rows
	rows, err = Db.Query(sentenceCount)
	defer rows.Close()

	if err != nil {
		return Resp, err
	}


	rows.Next()
	var regi sql.NullInt32
	err = rows.Scan(&regi)
	registers := int(regi.Int32)

	if err != nil {
		fmt.Println(err.Error())
		return Resp, err
	}

	if page > 0 {
		if registers > pageSize {
			limit = " LIMIT "+strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET "+strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	var orderBy string

	fmt.Println(orderField)
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Task_Id "
		case "T":
			orderBy = " ORDER BY task_title "
		case "D":
			orderBy = " ORDER BY task_description "
		case "F":
			orderBy = " ORDER BY is_done "	
		}
		if orderType == "D" {
			orderBy += " DESC"
		}
	}

	sentence += where + orderBy + limit

	fmt.Println("La sentencia es: "+sentence)

	rows, err = Db.Query(sentence)
	if err != nil {
		return Resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task
		var TaskId sql.NullInt32
		var TaskTitle sql.NullString
		var TaskDescription sql.NullString
		var TaskDone sql.NullInt32

		err := rows.Scan(&TaskId, &TaskTitle, &TaskDescription, &TaskDone)

		if err != nil {
			return Resp, err
		}

		t.TaskID = int(TaskId.Int32)
		t.TaskTitle = TaskTitle.String
		t.TaskDescription = TaskDescription.String
		t.TaskDone = int(TaskDone.Int32)
		task = append(task, t)

		fmt.Println("Finish")
	}

	Resp.TotalItems = registers
	Resp.Data = task

	return Resp, nil
}