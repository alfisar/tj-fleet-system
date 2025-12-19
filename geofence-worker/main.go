package main

import (
	"geofence-worker/config"
	"log"
)

func main() {
	poolData := config.DataPool.Get().(*config.Config)

	ch, err := poolData.ConnRabbit.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(
		"geofence_alerts",
		"worker geofence",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
