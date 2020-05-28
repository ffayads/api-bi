package models

import (
    "github.com/jinzhu/gorm"
)

type ReportsConclusions struct {
    gorm.Model
	Reports             Reports
	ReportsID           uint        `gorm:"foreignkey:ReportsID"`
	Description         string      `gorm:"not null"`
}
