package config

import (
	"fleet-ingestion/database"
	"fmt"
	"log"
	"os"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	_ = godotenv.Load(".env")
)

type Config struct {
	DBSql    *gorm.DB
	MQTTConn mqtt.Client
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
				DBSql:    DBSql,
				MQTTConn: InitConnMQTT(),
			}
		},
	}
	return &DataPool
}

func InitConnMQTT() (client mqtt.Client) {
	if os.Getenv("MQTTHOST") == "" {
		log.Fatal("Failed to connect to MQTT: MQTTHOST  is empty")
		return
	}

	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTTHOST")).
		SetClientID("fleet-subscriber-" + uuid.NewString())

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Failed to connect to MQTT: " + token.Error().Error())
		return
	}
	fmt.Println("Connected to Mosquitto")
	return
}
