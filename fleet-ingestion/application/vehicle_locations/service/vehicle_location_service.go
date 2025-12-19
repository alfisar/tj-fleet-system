package service

import (
	"context"
	"fleet-ingestion/application/vehicle_locations/repository"
	"fleet-ingestion/config"
	"fleet-ingestion/domain"
	"fleet-ingestion/helper/handler"
)

type vehicleLocationService struct {
	repo repository.VehicleLocationRepositoryContract
}

func NewVehicleLocationService(repo repository.VehicleLocationRepositoryContract) *vehicleLocationService {
	return &vehicleLocationService{
		repo: repo,
	}
}

func (s vehicleLocationService) VehicleLocation(ctx context.Context, poolData *config.Config, data domain.VehicleLocation) (err error) {
	defer handler.PanicError("vehicle_location", "VehicleLocation")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = s.repo.Insert(ctx, poolData.DBSql, data)
	return
}
