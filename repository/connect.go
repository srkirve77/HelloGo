package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "../../mysql"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:new-password@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to the database")
	con = db
	return db
}
