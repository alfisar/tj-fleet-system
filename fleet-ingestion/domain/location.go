package domain

type VehicleLocation struct {
	VihicleID string  `json:"vehicle_id" gorm:"column:vehicle_id"`
	Latitude  float64 `json:"lat" gorm:"column:latitude"`
	Langitude float64 `json:"lng" gorm:"column:langitude"`
	Timestamp int64   `json:"timestamp" gorm:"column:timestamp"`
}
