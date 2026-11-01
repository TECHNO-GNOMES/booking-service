package models

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Name string
}
