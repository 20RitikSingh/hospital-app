package handlers

import (
	"github.com/20ritiksingh/hospital-app/internal/service"
)

type APIHandler struct {
	authService    *service.AuthService
	patientService *service.PatientService
}

func NewAPIHandler(authService *service.AuthService, patientService *service.PatientService) *APIHandler {
	return &APIHandler{
		authService:    authService,
		patientService: patientService,
	}
}
