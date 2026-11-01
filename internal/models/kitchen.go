package models

import "gorm.io/gorm"

type Kitchen struct {
	gorm.Model
	Name string
}
