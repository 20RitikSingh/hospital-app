package mapper

import (
	"log"

	"github.com/20ritiksingh/hospital-app/internal/openapi"
	"github.com/20ritiksingh/hospital-app/models"
)

func MapSignupReqestToUser(r *openapi.SignupRequest) models.User {
	role, err := models.ParseRole(string(r.Role))
	if err != nil {
		log.Println("Error parsing role:", err)
		role = models.Receptionist // default to Receptionist if parsing fails
	}
	return models.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Role:     role,
	}
}
func MapLoginRequestToUser(r *openapi.LoginRequest) models.User {
	return models.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

func MapPatientToAPIPatient(p *models.Patient) openapi.Patient {
	return openapi.Patient{
		Id:     &p.ID,
		Name:   p.Name,
		Age:    p.Age,
		Email:  p.Email,
		Gender: openapi.PatientGender(p.Gender.String()),
	}
}

func MapApiNewPatientToPatient(p openapi.NewPatient) models.Patient {
	gender, err := models.ParseGender(string(p.Gender))
	if err != nil {
		gender = models.Other // default to Other if parsing fails
	}
	return models.Patient{
		Name:   p.Name,
		Age:    p.Age,
		Gender: gender,
		Email:  p.Email,
	}
}
