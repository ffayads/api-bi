package models

import (
    "github.com/jinzhu/gorm"
)

type Brands struct {
    gorm.Model
	Name                string
	Status              bool        `gorm:"not null"`
}
