package main

import (
	"encoding/json"
	"fmt"
	"mqtt-publisher/config"
	"mqtt-publisher/domain"
	"time"
)

func main() {
	client := config.InitConnMQTT()
	vehicleID := "B1234XYZ"
	topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		data := domain.VehicleLocation{
			VehicleID: vehicleID,
			Lat:       -6.200000 + float64(time.Now().Second())*0.0001,
			Lng:       106.816666 + float64(time.Now().Second())*0.0001,
			Timestamp: time.Now().Unix(),
		}

		payload, _ := json.Marshal(data)
		token := client.Publish(topic, 0, false, payload)
		token.Wait()
		fmt.Println("Published:", string(payload))
	}
}
