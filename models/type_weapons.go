package models

import (
    "github.com/jinzhu/gorm"
)

type TypeWeapons struct {
    gorm.Model
	Description         string      `gorm:"not null"`
	Status              bool        `gorm:"not null"`
}
