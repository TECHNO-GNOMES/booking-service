package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar    string
	FirstName string
	LastName  string
	City      string
	Birthday  *time.Time
	Gender    string
	Phone     string
	Address   string
	Email     string `gorm:"unique"`
	Password  string `gorm:"unique"`
	Role      string
	UpdatedAt *time.Time
}
