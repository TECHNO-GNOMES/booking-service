package models

import "gorm.io/gorm"

type InteriorStyle struct {
	gorm.Model
	Name string
}
