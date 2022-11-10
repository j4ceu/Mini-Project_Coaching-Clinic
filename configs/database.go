package configs

import (
	"Mini-Project_Coaching-Clinic/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	DB_URI := Cfg.DB_URI

	dsn := DB_URI // Build connection string
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	log.Println("Koneksi DB Berhasil")
	// initMigrate()

}

func initMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Coach{}, &models.Game{}, &models.CoachExperience{}, &models.CoachAvailability{}, &models.UserPayment{}, &models.UserBook{}) // Migrate the schema
	if err != nil {
		log.Fatal("Migration Gagal")
	} else {
		log.Println("Migration Berhasil")
	}
}
