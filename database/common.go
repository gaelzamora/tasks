package database

import (
	"database/sql"
	"fmt"
	"os"
	"tasks/models"
	"tasks/secrets"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel	, err = secrets.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	fmt.Println(Db)
	Db, err = sql.Open("mysql", ConStr(SecretModel))
	if err != nil {
		fmt.Print("There's an error")
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("There's other error")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conection Succesful of DB")
	return nil
}

func ConStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "tasks"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}