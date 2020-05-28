package models

import (
    "github.com/jinzhu/gorm"
)

type Users struct {
    gorm.Model
	Name                string
	Email               string      `gorm:"unique_index:idx_email_role"`
	Password            string      `gorm:"not null"`
	Status              bool        `gorm:"not null"`
	Phone               *string
}
