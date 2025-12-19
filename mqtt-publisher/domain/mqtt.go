package domain

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id"`
	Lat       float64 `json:"latitude"`
	Lng       float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}
