package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var (
	_ = godotenv.Load(".env")
)

type Config struct {
	ConnRabbit *amqp.Connection
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

	connRabbit := InitConfRabbit()

	DataPool := sync.Pool{
		New: func() interface{} {
			return &Config{
				ConnRabbit: connRabbit,
			}
		},
	}
	return &DataPool
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
