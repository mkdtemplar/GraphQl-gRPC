package database

import (
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connStr := "user=graphql password=graphql host=localhost port=5432 dbname=graphql_data sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("cannot connect to db %v", err)
	}
	DB = db

	log.Println("Connected to database")
}
