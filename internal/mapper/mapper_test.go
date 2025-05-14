package mapper

import (
	"testing"

	"github.com/20ritiksingh/hospital-app/internal/openapi"
	"github.com/20ritiksingh/hospital-app/models"
	"github.com/stretchr/testify/assert"
)

func TestMapPatientToAPIPatient(t *testing.T) {
	p := &models.Patient{
		ID:     123,
		Name:   "John Doe",
		Age:    30,
		Email:  "john@example.com",
		Gender: models.Male,
	}
	apiPatient := MapPatientToAPIPatient(p)
	assert.Equal(t, int64(123), int64(*apiPatient.Id))
	assert.Equal(t, "John Doe", apiPatient.Name)
	assert.Equal(t, 30, apiPatient.Age)
	assert.Equal(t, "john@example.com", apiPatient.Email)
	assert.Equal(t, openapi.PatientGender("male"), apiPatient.Gender)
}

func TestMapApiNewPatientToPatient(t *testing.T) {
	t.Run("valid gender", func(t *testing.T) {
		newP := openapi.NewPatient{
			Name:   "Jane",
			Age:    25,
			Gender: openapi.NewPatientGender("female"),
			Email:  "jane@example.com",
		}
		patient := MapApiNewPatientToPatient(newP)
		assert.Equal(t, "Jane", patient.Name)
		assert.Equal(t, 25, patient.Age)
		assert.Equal(t, models.Female, patient.Gender)
		assert.Equal(t, "jane@example.com", patient.Email)
	})

	t.Run("invalid gender defaults to Other", func(t *testing.T) {
		newP := openapi.NewPatient{
			Name:   "Alex",
			Age:    40,
			Gender: openapi.NewPatientGender("not-a-gender"),
			Email:  "alex@example.com",
		}
		patient := MapApiNewPatientToPatient(newP)
		assert.Equal(t, models.Other, patient.Gender)
	})
}
