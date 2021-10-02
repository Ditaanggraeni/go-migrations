package models

import (
	"gorm.io/gorm"
)

type KontrakMatkul struct {
	gorm.Model 
	MahasiswaID uint
	SemesterID uint
}