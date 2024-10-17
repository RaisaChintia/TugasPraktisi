package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/pemweb2"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database : ", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging to database : ", err)
	}
	fmt.Println("Successfully connected to MySQL database")
}
