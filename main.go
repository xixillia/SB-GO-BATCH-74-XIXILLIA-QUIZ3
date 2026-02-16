package main

import (
	"database/sql"
	"fmt"
	"quiz3/database"
	"quiz3/routers"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "go_quiz3"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)
	fmt.Println("Successfully connected to database!")
	defer DB.Close()

	router := routers.StartServer(DB).Run(":8080")
	if router != nil {
		panic(router)
	}
}
