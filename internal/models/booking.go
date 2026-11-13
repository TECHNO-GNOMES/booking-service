package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	UserID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID"`

	CheckIn  time.Time `gorm:"not null"`
	CheckOut time.Time `gorm:"not null"`

	TotalPrice int    `gorm:"not null"`
	Status     string `gorm:"not null;default:pending"`
	GuestCount int    `gorm:"not null;default:1"`
}
