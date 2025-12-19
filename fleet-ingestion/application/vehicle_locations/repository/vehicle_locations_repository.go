package repository

import (
	"context"
	"fleet-ingestion/domain"
	"fleet-ingestion/helper/consts"
	"fmt"

	"gorm.io/gorm"
)

type fleetIngestionRepositry struct{}

func NewFleetIngestionRepository() *fleetIngestionRepositry {
	return &fleetIngestionRepositry{}
}

func (r *fleetIngestionRepositry) Insert(ctx context.Context, conn *gorm.DB, data domain.VehicleLocation) (err error) {
	defer func() {
		if recov := recover(); recov != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", recov))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(fmt.Sprintf("Error Insert data in repository vihicle location : %s", consts.ErrMsgConnEmpty))
		return
	}

	err = conn.WithContext(ctx).Debug().Table("vehicle_location").Create(&data).Error
	return
}
