package validation

import (
	"fleet-ingestion/domain"
	"fleet-ingestion/helper/consts"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	Alphanumeric = validation.Match(regexp.MustCompile(consts.RegexAlphanumeric)).Error(consts.Alphanumeric)
	MaxChar9     = validation.Length(0, 9).Error(consts.MaxChar9)
	Required     = validation.Required.Error(consts.RequiredField)
)

func ValidationDataVehicleLocation(data domain.VehicleLocation) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.VehicleID, Required, Alphanumeric),
		validation.Field(&data.Latitude, Required),
		validation.Field(&data.Longitude, Required),
		validation.Field(&data.Timestamp, Required),
	)
	return
}
