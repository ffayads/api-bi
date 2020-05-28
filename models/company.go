package models

import (
    "github.com/jinzhu/gorm"
)

type Company struct {
    gorm.Model
	Name                string      `gorm:"not null"`
	Address             string
	Location            string
}

func (Company) TableName() string {
    return "company"
}
