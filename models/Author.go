package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Nama string
	Usia int
}
