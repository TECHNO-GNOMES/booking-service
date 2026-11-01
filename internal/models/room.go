package models

import (
	"time"

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

type CheckInAfter struct {
	gorm.Model
	TimeAfter *time.Time
}

type CheckInBefore struct {
	gorm.Model
	TimeBefore *time.Time
}

type CheckOutAfter struct {
	gorm.Model
	TimeAfter *time.Time
}

type CheckOutBefore struct {
	gorm.Model
	TimeBefore *time.Time
}
