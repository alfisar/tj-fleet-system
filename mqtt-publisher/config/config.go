package config

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load(".env")
)

func InitConnMQTT() (client mqtt.Client) {
	if os.Getenv("MQTTHOST") == "" || os.Getenv("CLIENTID") == "" {
		log.Fatal("Failed to connect to MQTT: MQTTHOST or CLIENTID is empty")
		return
	}

	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTTHOST")).
		SetClientID(os.Getenv("CLIENTID"))

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Failed to connect to MQTT: " + token.Error().Error())
		return
	}
	fmt.Println("Connected to Mosquitto")
	return
}
