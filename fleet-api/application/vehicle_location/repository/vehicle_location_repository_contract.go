package repository

import (
	"context"
	"fleet-api/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VehicleLocationRepositoryContract interface {
	GetLast(ctx context.Context, conn *gorm.DB, where any) (data domain.VehicleLocation, err error)
	GetHistory(ctx context.Context, conn *gorm.DB, where clause.Expr) (data []domain.VehicleLocation, err error)
}
