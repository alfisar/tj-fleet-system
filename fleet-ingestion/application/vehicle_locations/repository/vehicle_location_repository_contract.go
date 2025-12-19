package repository

import (
	"context"
	"fleet-ingestion/domain"

	"gorm.io/gorm"
)

type VehicleLocationRepositoryContract interface {
	Insert(ctx context.Context, conn *gorm.DB, data domain.VehicleLocation) (err error)
}
