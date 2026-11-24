package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UserID    uint   `gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserID"`
	Token     string `gorm:"uniqueIndex;not null"`
	Revoked   bool   `gorm:"default:false"`
	ExpiresAt *time.Time
}
