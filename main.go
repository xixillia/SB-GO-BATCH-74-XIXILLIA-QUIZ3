package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz3/database"
	"quiz3/routers"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// GANTI %d menjadi %s
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)
	fmt.Println("Successfully connected to database!")
	defer DB.Close()

	router := routers.StartServer(DB)
	router.Run(":" + os.Getenv("PORT"))
}
