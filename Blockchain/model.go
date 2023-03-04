package model

import "time"

// AVD : Autonomous Vehicle Data Block
type AVD struct {
	ID                      string        `json:"id"`
	LicensePlate            string        `json:"licensePlate"`
	VehicleModel            string        `json:"vehicleModel"`
	VinNumber               string        `json:"vinNumber"`
	DataTransmission        string        `json:"dataTransmission"`
	SurroundingVehicleData  []AVD         `json:"nearbyVehicles"`
}
