package domain

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id" gorm:"column:vehicle_id"`
	Latitude  float64 `json:"latitude" gorm:"column:latitude"`
	Longitude float64 `json:"longitude" gorm:"column:longitude"`
	Timestamp int64   `json:"timestamp" gorm:"column:timestamp"`
}

type Geofence struct {
	VehicleID string   `json:"vehicle_id"`
	Event     string   `json:"event"`
	Location  Location `json:"location"`
	Timestamp int64    `json:"timestamp"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude" `
}
