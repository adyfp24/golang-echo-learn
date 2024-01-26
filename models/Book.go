package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul    string
	Genre    string
	Harga    string
	AuthorID uint `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Author   Author
}
