package models

import (
	"gorm.io/gorm"
)

type Matkul struct {
	gorm.Model 
	Nama_matkul string `gorm:"type:varchar(50);not null" json:"nama_matkul"`
	Sks int`gorm:"type:int" json:"sks"`
}