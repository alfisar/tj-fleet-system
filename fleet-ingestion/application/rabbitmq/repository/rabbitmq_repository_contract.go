package rabbitmq

import "github.com/streadway/amqp"

type RabbitMQRepositoryContract interface {
	Publish(exchangeName string, routing string, message string, ch *amqp.Channel) (err error)
}
