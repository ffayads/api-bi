package models

import (
    "github.com/jinzhu/gorm"
)

type InstrumentsReports struct {
    gorm.Model
	Instruments         Instruments
	InstrumentsID       uint      `gorm:"primary_key;not null"`
	Reports             Reports
	ReportsID           uint      `gorm:"primary_key;not null"`
}
