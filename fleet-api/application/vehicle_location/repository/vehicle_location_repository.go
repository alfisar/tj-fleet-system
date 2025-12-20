package repository

import (
	"context"
	"fleet-api/domain"
	"fleet-api/helpers/errorhandler"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type vehicleLocationRepository struct{}

func NewVehicleLocationRepository() *vehicleLocationRepository {
	return &vehicleLocationRepository{}
}

func (r *vehicleLocationRepository) GetLast(ctx context.Context, conn *gorm.DB, where any) (data domain.VehicleLocation, err error) {
	defer func() {
		if recov := recover(); recov != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", recov))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(fmt.Sprintf("Error Get Last data in repository vehicle location : %s", errorhandler.ErrMsgConnEmpty))
		return
	}

	err = conn.WithContext(ctx).Debug().Table("vehicle_location").Where(where).Order("timestamp DESC").Find(&data).Error
	return
}

func (r *vehicleLocationRepository) GetHistory(ctx context.Context, conn *gorm.DB, where clause.Expr) (data []domain.VehicleLocation, err error) {
	defer func() {
		if recov := recover(); recov != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", recov))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(fmt.Sprintf("Error Get Hostory data in repository vehicle location : %s", errorhandler.ErrMsgConnEmpty))
		return
	}

	err = conn.WithContext(ctx).Debug().Table("vehicle_location").Where(where).Find(&data).Error
	return
}
