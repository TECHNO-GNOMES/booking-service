package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	PropertyID uint
	Property   Property `gorm:"foreignKey:PropertyID"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	Text   *string
	Rating int
}
