package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	database, err := sql.Open("postgres", "user=postgres password=Pass1234 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	row := database.QueryRow("INSERT INTO users (name, age) values ($1, $2)  RETURNING id", "AnuchitO", 19)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)
}
