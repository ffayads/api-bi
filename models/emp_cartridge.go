package models

import (
    "github.com/jinzhu/gorm"
)

type EmpCartridge struct {
    gorm.Model
	Emp                 Emp             `gorm:"foreignkey:EmpID"`
	EmpID               uint            `gorm:"not null"`
	Projectile          bool            `gorm:"not null"`
	Caliber             string          `gorm:"not null"`
	Quantity            int             `gorm:"not null"`
	Details             string          `gorm:"not null"`
	Shape               string          `gorm:"not null"`
	Size                string          `gorm:"not null"`
	EmpCartridgecol     string          `gorm:"not null"`
	CartridgeBrands     CartridgeBrands `gorm:"foreignkey:CartridgeBrandsID"`
	CartridgeBrandsID   uint            `gorm:"not null"`
	Observation         string          `gorm:"type:text"`
}

func (EmpCartridge) TableName() string {
    return "emp_cartridge"
}
