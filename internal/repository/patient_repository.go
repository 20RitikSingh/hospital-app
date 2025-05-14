package repository

import (
	"github.com/20ritiksingh/hospital-app/models"
	"gorm.io/gorm"
)

type PatientRepository interface {
	CreatePatient(patient *models.Patient) (*models.Patient, error)
	GetPatientByID(id uint) (*models.Patient, error)
	GetPatientByEmail(email string) (*models.Patient, error)
	UpdatePatient(id uint, patient *models.Patient) (*models.Patient, error)
	DeletePatient(id uint) error
	GetAllPatients() ([]*models.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	return patient, r.db.Create(patient).Error
}

func (r *patientRepository) GetPatientByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	return &patient, r.db.First(&patient, id).Error
}

func (r *patientRepository) GetPatientByEmail(email string) (*models.Patient, error) {
	var patient models.Patient
	return &patient, r.db.Where("email = ?", email).First(&patient).Error
}

func (r *patientRepository) UpdatePatient(id uint, patient *models.Patient) (*models.Patient, error) {
	var existingPatient models.Patient
	err := r.db.First(&existingPatient, id).Error
	if err != nil {
		return nil, err
	}

	// Update the patient record
	r.db.Model(&existingPatient).Updates(patient)
	return &existingPatient, nil
}
func (r *patientRepository) DeletePatient(id uint) error {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&patient).Error
}

func (r *patientRepository) GetAllPatients() ([]*models.Patient, error) {
	var patients []*models.Patient
	err := r.db.Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}
