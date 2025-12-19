package config

import (
	"fleet-ingestion/database"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	_ = godotenv.Load(".env")
)

type Config struct {
	DBSql      *gorm.DB
	MQTTConn   mqtt.Client
	ConnRabbit *amqp.Connection
	Coord      CoordTJ
}

type CoordTJ struct {
	Latitude  float64
	Longitude float64
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

	lat, err := strconv.ParseFloat(os.Getenv("LAT"), 64)
	if err != nil {
		log.Fatalln("LAT is empty")
	}

	lon, err := strconv.ParseFloat(os.Getenv("LON"), 64)
	if err != nil {
		log.Fatalln("LAT is empty")
	}

	connMQTT := InitConnMQTT()
	connRabbit := InitConfRabbit()

	DataPool := sync.Pool{
		New: func() interface{} {
			return &Config{
				DBSql:      DBSql,
				MQTTConn:   connMQTT,
				ConnRabbit: connRabbit,
				Coord: CoordTJ{
					Latitude:  lat,
					Longitude: lon,
				},
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

func InitConfRabbit() (conn *amqp.Connection) {
	fmt.Println("Starting rabbitmq")

	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort := os.Getenv("AMQP_PORT")

	if amqpUser == "" || amqpPass == "" || amqpHost == "" || amqpPort == "" {
		log.Fatalln(fmt.Errorf("Failed Connect Rabbit : Invalid Data Rabbit"))
	}
	conRabbit := fmt.Sprintf("amqp://%s:%s@%s:%s/", amqpUser, amqpPass, amqpHost, amqpPort)
	conn, err := amqp.Dial(conRabbit)
	if err != nil {
		log.Fatalf("Gagal terkoneksi ke RabbitMQ: %v", err)
	}

	fmt.Println("Successfullt connect rabbitmq")
	return
}
