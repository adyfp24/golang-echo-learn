package models

import (
	"time"
)

type Author struct {
	ID        uint 		`gorm:"primaryKey"`
	Nama      string
	Usia      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
