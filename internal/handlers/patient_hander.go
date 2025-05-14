package handlers

import (
	"strconv"

	"github.com/20ritiksingh/hospital-app/internal/mapper"
	"github.com/20ritiksingh/hospital-app/internal/openapi"
	"github.com/gin-gonic/gin"
)

func (h *APIHandler) GetAllPatients(ctx *gin.Context) {
	patients, err := h.patientService.ListPatients()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	var apiPatients []openapi.Patient
	for _, patient := range patients {
		apiPatient := mapper.MapPatientToAPIPatient(patient)
		apiPatients = append(apiPatients, apiPatient)
	}
	ctx.JSON(200, apiPatients)

}
func (h *APIHandler) CreatePatient(ctx *gin.Context) {
	var req openapi.NewPatient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	patient := mapper.MapApiNewPatientToPatient(req)
	createdPatient, err := h.patientService.CreatePatient(patient)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	apiPatient := mapper.MapPatientToAPIPatient(createdPatient)
	ctx.JSON(201, apiPatient)
}

func (h *APIHandler) GetPatientByID(ctx *gin.Context) {
	id := ctx.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}
	patient, err := h.patientService.GetPatientByID(uint(patientID))
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Patient not found"})
		return
	}
	apiPatient := mapper.MapPatientToAPIPatient(patient)
	ctx.JSON(200, apiPatient)
}
func (h *APIHandler) UpdatePatientByID(ctx *gin.Context) {
	id := ctx.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	var req openapi.NewPatient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	patient := mapper.MapApiNewPatientToPatient(req)
	updatedPatient, err := h.patientService.UpdatePatient(uint(patientID), patient)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	apiPatient := mapper.MapPatientToAPIPatient(updatedPatient)
	ctx.JSON(200, apiPatient)
}
func (h *APIHandler) DeletePatientByID(ctx *gin.Context) {
	id := ctx.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	err = h.patientService.DeletePatient(uint(patientID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(204, nil)
}
