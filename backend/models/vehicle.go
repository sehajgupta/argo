package models

import "time"

type Vehicle struct {
	ID           string    `db:"id" json:"id"`
	DriverName   string    `db:"driver_name" json:"driverName"`
	LicensePlate string    `db:"license_plate" json:"licensePlate"`
	ETA          time.Time `db:"eta" json:"eta"`
	Location     string    `db:"location" json:"location"`
}
