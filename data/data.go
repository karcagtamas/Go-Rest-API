package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/karcatamas/restapi/models"
)

func Connect() {
	db, err := sql.Open("mysql", "gitbucket:gitbucket@tcp(127.0.0.1:3306)/csomormaker")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	selectQ, _ := db.Query("SELECT * FROM roles;")

	defer selectQ.Close()

	for selectQ.Next() {
		var result models.Role
		selectQ.Scan(&result.ID, &result.Name, &result.AccessLevel)
		fmt.Println(result)
	}
}
