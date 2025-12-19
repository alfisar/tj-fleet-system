package service

import (
	"context"
	"fleet-ingestion/config"
	"fleet-ingestion/domain"
)

type VehicleLocationServiceContract interface {
	VehicleLocation(ctx context.Context, poolData *config.Config, data domain.VehicleLocation) (err error)
}
