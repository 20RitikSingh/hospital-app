package database

import (
	"log"

	"github.com/20ritiksingh/hospital-app/internal/utils"
	"github.com/20ritiksingh/hospital-app/models"
	"gorm.io/gorm"
)

// SeedData seeds the database with initial data
func SeedData(db *gorm.DB) error {
	// Seed Receptionists
	receptionists := []models.User{
		{Name: "Alice", Email: "alice@mail.com", Password: "receptionist123", Role: models.Receptionist},
		{Name: "Bob", Email: "bob@mail.com", Password: "receptionist123", Role: models.Receptionist},
		{Name: "Charlie", Email: "charlie@mail.com", Password: "receptionist123", Role: models.Receptionist},
	}
	for _, r := range receptionists {
		hashedPassword, err := utils.HashPassword(r.Password)
		if err != nil {
			log.Println("Error hashing password:", err)
			continue
		}
		r.Password = string(hashedPassword)
		if err := db.Create(&r).Error; err != nil {
			log.Println("Error seeding receptionist:", err)
		}
	}

	// Seed Doctors
	doctors := []models.User{
		{Name: "Dr. Smith", Email: "smith@mail.com", Password: "doctor123", Role: models.Doctor},
		{Name: "Dr. Jones", Email: "jones@mail.com", Password: "doctor123", Role: models.Doctor},
		{Name: "Dr. Brown", Email: "brown@mail.com", Password: "doctor123", Role: models.Doctor},
	}
	for _, d := range doctors {
		hashedPassword, err := utils.HashPassword(d.Password)
		if err != nil {
			log.Println("Error hashing password:", err)
			continue
		}
		d.Password = string(hashedPassword)
		if err := db.Create(&d).Error; err != nil {
			log.Println("Error seeding doctor:", err)
		}
	}

	// Seed Patients
	patients := []models.Patient{
		{Name: "John Doe", Age: 45, Gender: models.Male, Email: "john@mail.com"},
		{Name: "Jane Smith", Age: 36, Gender: models.Female, Email: "jane@mail.com"},
	}
	for _, p := range patients {
		if err := db.Create(&p).Error; err != nil {
			return err
		}
	}

	return nil
}
