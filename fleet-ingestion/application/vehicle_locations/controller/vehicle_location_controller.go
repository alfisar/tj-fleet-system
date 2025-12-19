package controller

import (
	"context"
	"fleet-ingestion/application/vehicle_locations/service"
	"fleet-ingestion/config"
	"fleet-ingestion/domain"
	"fleet-ingestion/helper/validation"
)

type vehicleLocationController struct {
	serv service.VehicleLocationServiceContract
}

func NewVehicleLocationController(serv service.VehicleLocationServiceContract) *vehicleLocationController {
	return &vehicleLocationController{
		serv: serv,
	}
}

func (c *vehicleLocationController) InitPoolData() *config.Config {
	poolData := config.DataPool.Get().(*config.Config)
	return poolData
}

func (c vehicleLocationController) VehicleLocation(ctx context.Context, data domain.VehicleLocation) (err error) {
	poolData := c.InitPoolData()

	err = validation.ValidationDataVehicleLocation(data)
	if err != nil {
		return
	}

	err = c.serv.VehicleLocation(ctx, poolData, data)
	return
}
