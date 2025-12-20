package surounding

import (
	"log"

	"github.com/streadway/amqp"
)

func InitExchangeQueue(conn *amqp.Connection, exchangeName string, queueName string, key string) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("Error setup channel rabbitMQ : " + err.Error())
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln("Error Declare Exchange : " + err.Error())
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln("Error Declare Queue : " + err.Error())
	}

	err = ch.QueueBind(
		q.Name,
		key,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("Error Bind Queue and exchange : " + err.Error())
	}
}
