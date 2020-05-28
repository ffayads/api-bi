package models

import (
    "github.com/jinzhu/gorm"
)

type EmpImage struct {
    gorm.Model
	Emp                 Emp         `gorm:"foreignkey:EmpID"`
	EmpID               uint        `gorm:"not null"`
	UrlImage            string      `gorm:"not null"`
	Status              bool        `gorm:"not null"`
}

func (EmpImage) TableName() string {
    return "emp_image"
}
