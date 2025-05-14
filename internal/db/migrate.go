package database

import (
	"github.com/20ritiksingh/hospital-app/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(DB *gorm.DB) error {
	return DB.AutoMigrate(&models.User{}, &models.Patient{}, &models.MedicalRecord{})
}
