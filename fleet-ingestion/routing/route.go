package routing

import (
	"context"
	"encoding/json"
	"fleet-ingestion/domain"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Routing(message mqtt.Message) {
	controllerVehicle := InitVehicle()
	data := domain.VehicleLocation{}

	err := json.Unmarshal(message.Payload(), &data)
	if err != nil {
		fmt.Println("Error parsing data with message ID " + string(message.MessageID()) + " And data is " + string(message.Payload()))
	}

	ctx := context.Background()
	err = controllerVehicle.VehicleLocation(ctx, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
