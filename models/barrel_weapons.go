package models

import (
    "github.com/jinzhu/gorm"
)

type BarrelWeapons struct {
    gorm.Model
	Description         string      `gorm:"not null"`
	Status              bool        `gorm:"not null"`
}
