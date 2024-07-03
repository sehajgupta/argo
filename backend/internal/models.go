package server

import "time"

// Trip represents a trip in the system
type Trip struct {
	ID            int        `json:"id"`
	UserID        int        `json:"user_id"`
	StartLocation string     `json:"start_location"`
	EndLocation   string     `json:"end_location"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	DriverInfo    string     `json:"driver_info"`
	LicensePlate  string     `json:"license_plate"`
	Status        string     `json:"status"`
}

// Vehicle represents a vehicle in the system
type Vehicle struct {
	ID              int       `json:"id"`
	TripID          int       `json:"trip_id"`
	CurrentLocation string    `json:"current_location"`
	ETA             time.Time `json:"eta"`
	Status          string    `json:"status"`
	DriverInfo      string    `json:"driver_info"`
	LicensePlate    string    `json:"license_plate"`
}
