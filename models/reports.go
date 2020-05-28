package models

import (
    "github.com/jinzhu/gorm"
)

type Reports struct {
    gorm.Model
	Company             Company         `gorm:"foreignkey:CompanyID"`
	CompanyID           uint
	Users               Users           `gorm:"foreignkey:UsersID"`
	UsersID             uint
	Reason              string          `gorm:"type:text"`
	Name                string
    Instruments         []Instruments   `gorm:"many2many:instruments_reports;"`
	Status              bool            `gorm:"not null"`
}
