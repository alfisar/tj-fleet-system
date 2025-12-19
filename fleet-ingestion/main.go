package main

import (
	"fleet-ingestion/config"
	"fleet-ingestion/routing"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	poolData := config.DataPool.Get().(*config.Config)
	messages := make(chan mqtt.Message, 100)

	topic := "/fleet/vehicle/+/location"

	if token := poolData.MQTTConn.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println("get message from MQTT")
		messages <- msg
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Subscribe error: %v", token.Error())
	}

	workerCount, err := strconv.Atoi(os.Getenv("worker"))
	if err != nil {
		log.Fatalln("Worker must be numeric")
	}

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for msg := range messages {
				fmt.Println("worker " + strconv.Itoa(workerID) + " get message : " + string(msg.Payload()))
				routing.Routing(msg)
			}
		}(i)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Consumer succefully running")
	<-signals
	fmt.Println("Received shutdown signal")

	close(messages)

	wg.Wait()
	fmt.Println("All workers have shut down")

}
