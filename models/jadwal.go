package models

import (
	"gorm.io/gorm"
)

type Jadwal struct {
	gorm.Model 
	Jadwal string `gorm:"type:varchar(200);not null" json:"jadwal"`
	Matakuliah_id uint
}