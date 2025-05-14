package models

import (
	"fmt"
	"time"
)

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

type Patient struct {
	ID             int             `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"type:varchar(100);not null"`
	Age            int             `gorm:"not null;check:age >= 0"`
	Email          string          `gorm:"type:varchar(100);uniqueIndex;not null"`           // optional: prevent negative age
	Gender         Gender          `gorm:"type:int;not null"`                                // define possible values
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE"` // delete records if patient is deleted
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime"`
}

// String method to convert Gender to string
func (r Gender) String() string {
	switch r {
	case Male:
		return "male"
	case Female:
		return "female"
	case Other:
		return "other"
	default:
		return "unknown"
	}
}

// Scan method for converting the database value to Gender
func (r *Gender) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*r = Gender(v)
		return nil
	case int:
		*r = Gender(v)
		return nil
	default:
		return fmt.Errorf("cannot scan value %v into Gender", value)
	}
}

// Value method for converting UserRole to database value (int)
func (r Gender) Value() (int, error) {
	return int(r), nil
}

// ParseGender function to convert string
func ParseGender(s string) (Gender, error) {
	switch s {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	case "other":
		return Other, nil
	default:
		return -1, fmt.Errorf("invalid gender: %s", s)
	}
}
