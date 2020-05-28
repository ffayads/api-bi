package models

import (
    "github.com/jinzhu/gorm"
)

type Instruments struct {
    gorm.Model
	Description     string
	ModelInstrument string          `gorm:"not null"`
	Serial          string          `gorm:"not null"`
	Brands          Brands          `gorm:"foreignkey:BrandsID"`
    BrandsID        uint            `gorm:"not null"`
	Status          bool            `gorm:"not null"`
}
