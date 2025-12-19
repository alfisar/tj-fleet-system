package routing

import (
	"fleet-api/application/vehicle_location/controller"

	"github.com/gofiber/fiber/v2"
)

type VehicleRoute struct {
	controll controller.VehicleLocationControllerContract
}

func NewVehicleRoute(controll controller.VehicleLocationControllerContract) *VehicleRoute {
	return &VehicleRoute{
		controll: controll,
	}
}

func (r *VehicleRoute) vehicleRoute(v1 fiber.Router) {
	v1.Get("/vehicles/:vehicleID/location", r.controll.GetLast)
	v1.Get("/vehicles/:vehicleID/history", r.controll.GetHistory)
}
