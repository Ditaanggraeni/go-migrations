package models

import (
	"gorm.io/gorm"
)

type Mahasiswa struct {
	gorm.Model 
	Nama string `gorm:"type:varchar(200);not null" json:"nama"`
	Nim int`gorm:"size:20" json:"nim"`
	Jurusan string `gorm:"type:varchar(200);not null" json:"jurusan"`
}