package config

import (
	"fleet-api/database"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	_ = godotenv.Load(".env")
)

type Config struct {
	DBSql *gorm.DB
}

var (
	_                  = godotenv.Load(".env")
	DataPool sync.Pool = *InitConfigPool()
)

func InitConfigPool() *sync.Pool {
	var DBSql *gorm.DB
	fmt.Println("DB_USE : " + os.Getenv("DB_USE"))
	switch os.Getenv("DB_USE") {
	case
		"postgress":
		DBSql = database.NewDatabaseMySql()

	default:
		log.Fatalln("error DB Use")
	}

	DataPool := sync.Pool{
		New: func() interface{} {
			return &Config{
				DBSql: DBSql,
			}
		},
	}
	return &DataPool
}
