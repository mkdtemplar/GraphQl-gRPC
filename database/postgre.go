package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=graphql password=graphql host=localhost port=5432 dbname=graphql_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("cannot connect to db %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db

	log.Println("Connected to database")
}
