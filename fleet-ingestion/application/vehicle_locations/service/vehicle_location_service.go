package service

import (
	"context"
	"encoding/json"
	repositoryRabbit "fleet-ingestion/application/rabbitmq/repository"
	"fleet-ingestion/application/vehicle_locations/repository"
	"fleet-ingestion/config"
	"fleet-ingestion/domain"
	"fleet-ingestion/helper/handler"
	"fmt"
)

type vehicleLocationService struct {
	repo       repository.VehicleLocationRepositoryContract
	repoRabbit repositoryRabbit.RabbitMQRepositoryContract
}

func NewVehicleLocationService(repo repository.VehicleLocationRepositoryContract, repoRabbit repositoryRabbit.RabbitMQRepositoryContract) *vehicleLocationService {
	return &vehicleLocationService{
		repo:       repo,
		repoRabbit: repoRabbit,
	}
}

func (s vehicleLocationService) VehicleLocation(ctx context.Context, poolData *config.Config, data domain.VehicleLocation) (err error) {
	defer handler.PanicError("vehicle_location", "VehicleLocation")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = s.repo.Insert(ctx, poolData.DBSql, data)
	if err != nil {
		fmt.Println(err)
	}

	radiusData := isRadius(data.Latitude, data.Longitude, poolData.Coord.Latitude, poolData.Coord.Longitude)

	if radiusData {
		data := domain.Geofence{
			VehicleID: data.VehicleID,
			Event:     "geofence_entry",
			Location: domain.Location{
				Latitude:  data.Latitude,
				Longitude: data.Longitude,
			},
			Timestamp: data.Timestamp,
		}

		message, errData := json.Marshal(&data)
		if errData != nil {
			err = errData
			return
		}

		channel, errData := poolData.Rabbit.ConnRabbit.Channel()
		if errData != nil {
			err = errData
			return
		}
		defer channel.Close()

		err = s.repoRabbit.Publish(poolData.Rabbit.ExchangeName, poolData.Rabbit.Key, string(message), channel)
		if err != nil {
			return
		}
	}
	return
}
