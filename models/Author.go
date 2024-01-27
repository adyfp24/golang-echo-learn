package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Nama string
	Usia int
}
