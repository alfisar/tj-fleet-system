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

	// membuat file migration
	// migrate create -ext sql -dir database/migrations <nama filenya>

	// migrasi semua file migration
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations up

	// rollback semua file migration
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations down

	// migrasi 1 version
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations up 1

	// rollback 1 version
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations down 1

	// check version migrate
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations version

	// force update migrate yang dirty
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations force <angka dari namanya >
}
