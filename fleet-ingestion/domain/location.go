package domain

type VehicleLocation struct {
	VihicleID string  `json:"vehicle_id" gorm:"column:vehicle_id"`
	Latitude  float64 `json:"latitude" gorm:"column:latitude"`
	Langitude float64 `json:"longitude" gorm:"column:longitude"`
	Timestamp int64   `json:"timestamp" gorm:"column:timestamp"`
}
