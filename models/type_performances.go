package models

import (
    "github.com/jinzhu/gorm"
)

type TypePerformances struct {
    gorm.Model
	Description         string      `gorm:"not null"`
	Status              bool        `gorm:"not null"`
}
