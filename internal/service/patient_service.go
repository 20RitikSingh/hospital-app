package service

import (
	"errors"

	"github.com/20ritiksingh/hospital-app/internal/repository"
	"github.com/20ritiksingh/hospital-app/models"
)

type PatientService struct {
	repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

// CreatePatient creates a new patient record
func (s *PatientService) CreatePatient(patient models.Patient) (*models.Patient, error) {
	// Check if patient already exists (based on some unique field like email)
	existingPatient, err := s.repo.GetPatientByEmail(patient.Email)
	if err == nil && existingPatient != nil {
		return nil, errors.New("patient with this email already exists")
	}

	// Save the new patient to the database
	createdPatient, err := s.repo.CreatePatient(&patient)
	if err != nil {
		return nil, err
	}

	return createdPatient, nil
}

// GetPatientByID retrieves a patient by their ID
func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
	patient, err := s.repo.GetPatientByID(id)
	if err != nil {
		return nil, errors.New("patient not found")
	}
	return patient, nil
}

// UpdatePatient updates an existing patient record
func (s *PatientService) UpdatePatient(id uint, patient models.Patient) (*models.Patient, error) {
	// Check if patient exists
	_, err := s.repo.GetPatientByID(id)
	if err != nil {
		return nil, errors.New("patient not found")
	}

	// Save the updated patient
	updatedPatient, err := s.repo.UpdatePatient(id, &patient)
	if err != nil {
		return nil, err
	}

	return updatedPatient, nil
}

// DeletePatient removes a patient record
func (s *PatientService) DeletePatient(id uint) error {
	// Check if patient exists
	_, err := s.repo.GetPatientByID(id)
	if err != nil {
		return errors.New("patient not found")
	}

	// Delete the patient record
	err = s.repo.DeletePatient(id)
	if err != nil {
		return err
	}

	return nil
}

// ListPatients retrieves all patients
func (s *PatientService) ListPatients() ([]*models.Patient, error) {
	patients, err := s.repo.GetAllPatients()
	if err != nil {
		return nil, err
	}
	return patients, nil
}
