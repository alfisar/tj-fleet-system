package rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type rabbitMQRepository struct{}

func NewRabbitMQ() *rabbitMQRepository {
	return &rabbitMQRepository{}
}

func (r *rabbitMQRepository) Publish(exchangeName string, routing string, message string, ch *amqp.Channel) (err error) {

	msg := amqp.Publishing{
		Body: []byte(message + time.Now().String()),
	}

	// Mengirim pesan ke exchange
	err = ch.Publish(
		exchangeName,
		routing,
		false,
		false,
		msg,
	)
	if err != nil {
		log.Printf("Gagal mengirim pesan: %v", err)
	}

	return
}
