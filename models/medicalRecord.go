package models

import (
	"time"
)

type MedicalRecord struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`                            // Explicit auto increment for clarity
	PatientID   uint      `gorm:"not null"`                                            // Foreign key for patient
	Patient     Patient   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"`    // Ensuring cascading delete for Patient
	Symptoms    string    `gorm:"type:text"`                                           // Text field for symptoms
	Diagnosis   string    `gorm:"type:text"`                                           // Text field for diagnosis
	Treatment   string    `gorm:"type:text"`                                           // Text field for treatment
	UpdatedByID uint      `gorm:"not null"`                                            // Foreign key for doctor
	UpdatedBy   User      `gorm:"foreignKey:UpdatedByID;constraint:OnDelete:SET NULL"` // Set null if the doctor is deleted
	CreatedAt   time.Time `gorm:"autoCreateTime"`                                      // Automatically set creation time
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`                                      // Automatically update the time
}
