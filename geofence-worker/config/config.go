package config

import (
	"fmt"
	"geofence-worker/helper/surounding"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var (
	_ = godotenv.Load(".env")
)

type Config struct {
	Rabbit RabbitMQ
}
type RabbitMQ struct {
	ConnRabbit   *amqp.Connection
	ExchangeName string
	QueueName    string
	Key          string
}

var (
	_                  = godotenv.Load(".env")
	DataPool sync.Pool = *InitConfigPool()
)

func InitConfigPool() *sync.Pool {

	connRabbit, exchange, queue, key := InitConfRabbit()

	DataPool := sync.Pool{
		New: func() interface{} {
			return &Config{
				Rabbit: RabbitMQ{
					ConnRabbit:   connRabbit,
					ExchangeName: exchange,
					QueueName:    queue,
					Key:          key,
				},
			}
		},
	}
	return &DataPool
}

func InitConfRabbit() (conn *amqp.Connection, exchangeName string, queue string, key string) {
	fmt.Println("Starting rabbitmq")

	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort := os.Getenv("AMQP_PORT")
	amqpExchange := os.Getenv("AMQP_EXCHANGE")
	amqpQueue := os.Getenv("AMQP_QUEUE")
	amqpKey := os.Getenv("AMQP_KEY")

	if amqpUser == "" || amqpPass == "" || amqpHost == "" || amqpPort == "" || amqpExchange == "" || amqpQueue == "" || amqpKey == "" {
		log.Fatalln(fmt.Errorf("Failed Connect Rabbit : Invalid Data Rabbit"))
	}
	conRabbit := fmt.Sprintf("amqp://%s:%s@%s:%s/", amqpUser, amqpPass, amqpHost, amqpPort)
	conn, err := amqp.Dial(conRabbit)
	if err != nil {
		for i := 0; i < 5; i++ {
			conRabbit := fmt.Sprintf("amqp://%s:%s@%s:%s/", amqpUser, amqpPass, amqpHost, amqpPort)
			conn, err = amqp.Dial(conRabbit)
			if err != nil {
				fmt.Println("Waiting for RabbitMQ...")
				time.Sleep(2 * time.Second)
			} else {
				break
			}
		}
		if err != nil {
			log.Fatalf("Gagal terkoneksi ke RabbitMQ: %v", err)
		}
	}

	fmt.Println("Successfullt connect rabbitmq")

	surounding.InitExchangeQueue(conn, amqpExchange, amqpQueue, amqpKey)
	exchangeName = amqpExchange
	key = amqpKey
	queue = amqpQueue
	return
}
