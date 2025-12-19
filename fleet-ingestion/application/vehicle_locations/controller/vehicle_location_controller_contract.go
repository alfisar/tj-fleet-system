package controller

import (
	"context"
	"fleet-ingestion/domain"
)

type VehicleLocationControllerContract interface {
	VehicleLocation(ctx context.Context, data domain.VehicleLocation) (err error)
}
