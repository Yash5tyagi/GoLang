package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost:3306)/todo_api")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
