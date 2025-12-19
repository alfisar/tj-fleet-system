package routing

import (
	"fleet-api/application/vehicle_location/controller"
	"fleet-api/application/vehicle_location/repository"
	"fleet-api/application/vehicle_location/service"
)

func VehicleInit() *VehicleRoute {
	repo := repository.NewVehicleLocationRepository()
	serv := service.NewVehicleLocationService(repo)
	control := controller.NewVehicleLocationController(serv)
	return NewVehicleRoute(control)
}
