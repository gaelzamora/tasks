package models

type SecretRDSJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine string `json:"engine"`
	Host string `json:"host"`
	Port int `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type Task struct {
	TaskID int `json:"taskID"`
	TaskTitle string `json:"taskTitle"`
	TaskDescription string `json:"taskDescription"`
	TaskDone int `json:"taskDone"`
}