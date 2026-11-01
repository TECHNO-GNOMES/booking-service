package models

import "gorm.io/gorm"

type Parking struct {
	gorm.Model
	Selection bool
}

type ParkingSpaces struct {
	gorm.Model
	Count int
}
