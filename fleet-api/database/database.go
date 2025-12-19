package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseMySql() *gorm.DB {
	fmt.Println("starting DB....")
	dbHost := os.Getenv("DB_SQL_HOST")
	dbPort := os.Getenv("DB_SQL_PORT")
	dbUser := os.Getenv("DB_SQL_USER")
	dbPass := os.Getenv("DB_SQL_PSWD")
	dbName := os.Getenv("DB_SQL_NAME")

	if dbHost == "" || dbName == "" || dbPort == "" || dbUser == "" {
		log.Println("Postgress Database Connection failed...")
		log.Fatalln("invalid data Postgress")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Println("Postgress Database Connection failed...")
		log.Fatalln(err)
	}

	fmt.Println("Successfully connect to DB :) ")

	return db

}
