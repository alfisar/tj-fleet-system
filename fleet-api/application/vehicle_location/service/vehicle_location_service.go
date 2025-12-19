package service

import (
	"context"
	"fleet-api/application/vehicle_location/repository"
	"fleet-api/config"
	"fleet-api/domain"
	"fleet-api/helpers/errorhandler"
	"fleet-api/helpers/handler"

	"gorm.io/gorm"
)

type vehicleLocationService struct {
	repo repository.VehicleLocationRepositoryContract
}

func NewVehicleLocationService(repo repository.VehicleLocationRepositoryContract) *vehicleLocationService {
	return &vehicleLocationService{
		repo: repo,
	}
}

func (s *vehicleLocationService) GetLast(ctx context.Context, poolData *config.Config, vehicleID string) (data domain.VehicleLocation, err domain.ErrorData) {
	var errData error
	defer handler.PanicError("Vehicle_location", "GetLast")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	where := map[string]any{
		"vehicle_id": vehicleID,
	}
	data, errData = s.repo.GetLast(ctx, poolData.DBSql, where)
	if errData != nil {
		err = errorhandler.ErrRecordNotFound()
	}
	return

}

func (s *vehicleLocationService) GetHistory(ctx context.Context, poolData *config.Config, vehicleID string, startDate int64, endDate int64) (data []domain.VehicleLocation, err domain.ErrorData) {
	var errData error
	defer handler.PanicError("Vehicle_location", "GetLast")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	data = make([]domain.VehicleLocation, 0)
	where := gorm.Expr("vehicle_id = ? AND timestamp BETWEEN ? AND ? ", vehicleID, startDate, endDate)
	data, errData = s.repo.GetHistory(ctx, poolData.DBSql, where)
	if errData != nil {
		err = errorhandler.ErrGetData(errData)
	}
	return

}
