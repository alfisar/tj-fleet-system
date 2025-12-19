package service

import (
	"context"
	"fleet-api/config"
	"fleet-api/domain"
)

type VehicleLocationServiceContract interface {
	GetLast(ctx context.Context, poolData *config.Config, vehicleID string) (data domain.VehicleLocation, err domain.ErrorData)
	GetHistory(ctx context.Context, poolData *config.Config, vehicleID string, startDate int64, endDate int64) (data []domain.VehicleLocation, err domain.ErrorData)
}
