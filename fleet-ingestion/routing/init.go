package routing

import (
	repositoryRabbit "fleet-ingestion/application/rabbitmq/repository"
	"fleet-ingestion/application/vehicle_locations/controller"
	"fleet-ingestion/application/vehicle_locations/repository"
	"fleet-ingestion/application/vehicle_locations/service"
)

func InitVehicle() controller.VehicleLocationControllerContract {
	repo := repository.NewFleetIngestionRepository()
	repoRabbit := repositoryRabbit.NewRabbitMQ()
	serv := service.NewVehicleLocationService(repo, repoRabbit)
	controller := controller.NewVehicleLocationController(serv)
	return controller
}
