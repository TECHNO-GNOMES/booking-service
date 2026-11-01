package models

import "gorm.io/gorm"

type LocationType struct {
	gorm.Model
	Name string
}
