package routing

import (
	"fleet-ingestion/application/vehicle_locations/controller"
	"fleet-ingestion/application/vehicle_locations/repository"
	"fleet-ingestion/application/vehicle_locations/service"
)

func InitVehicle() controller.VehicleLocationControllerContract {
	repo := repository.NewFleetIngestionRepository()
	serv := service.NewVehicleLocationService(repo)
	controller := controller.NewVehicleLocationController(serv)
	return controller
}
