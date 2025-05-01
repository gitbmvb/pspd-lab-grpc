package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB * sql.DB

func Init() {
	var err error
	connStr := "host=localhost port=5432 user=admin password=admin dbname=pspdlabs sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	
	log.Println("Connected to database")
}