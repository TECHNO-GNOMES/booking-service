package models

import "gorm.io/gorm"

type Features struct {
	gorm.Model
	Name string
}
