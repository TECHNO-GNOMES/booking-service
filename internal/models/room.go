package models

import (
	"gorm.io/gorm"
)

type RoomCategory struct {
	gorm.Model
	Name string
}

type SleepingPlaces struct {
	gorm.Model
	Count int
}

type AmenitiesRoom struct {
	gorm.Model
	Name string
}

type ServicesRoom struct {
	gorm.Model
	Name string
}
