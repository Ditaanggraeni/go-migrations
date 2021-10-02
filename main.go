package main

import (
	"go-migrations/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=29dita dbname=absen port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(
		&models.Mahasiswa{},
		&models.Matkul{},
		&models.Jadwal{},
		&models.Semester{},
		&models.KontrakMatkul{},
	)
}