package models

import (
	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model 
	Semester string `gorm:"type:int;not null" json:"semester"`
}