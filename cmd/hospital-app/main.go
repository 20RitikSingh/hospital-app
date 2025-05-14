package main

import (
	"log"
	"net/http"

	"github.com/20ritiksingh/hospital-app/internal/config"
	database "github.com/20ritiksingh/hospital-app/internal/db"
	"github.com/20ritiksingh/hospital-app/internal/handlers"
	"github.com/20ritiksingh/hospital-app/internal/middleware"
	"github.com/20ritiksingh/hospital-app/internal/repository"
	"github.com/20ritiksingh/hospital-app/internal/service"
	"github.com/20ritiksingh/hospital-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	// Initialize the database connection

	// Initialize the repositories
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else {
		log.Println("Connected to database")
	}

	// migrate the database
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	} else {
		log.Println("Migration successful")
	}

	// Seed the database
	err = database.SeedData(db)
	if err != nil {
		log.Fatalf("could not seed db: %v", err)
	}

	// Initialize the repositories
	authRepo := repository.NewAuthRepository(db)
	patientRepo := repository.NewPatientRepository(db)

	// Initialize the services
	authService := service.NewAuthService(authRepo)
	patientService := service.NewPatientService(patientRepo)

	handler := handlers.NewAPIHandler(authService, patientService)

	r := gin.Default()

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)

	api := r.Group("/patients")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		api.GET("", middleware.RoleMiddleware(models.Receptionist.String(), models.Doctor.String()), handler.GetAllPatients)
		api.POST("", middleware.RoleMiddleware(models.Receptionist.String()), handler.CreatePatient)
		api.GET(":id", middleware.RoleMiddleware(models.Receptionist.String(), models.Doctor.String()), handler.GetPatientByID)
		api.PUT(":id", middleware.RoleMiddleware(models.Receptionist.String(), models.Doctor.String()), handler.UpdatePatientByID)
		api.DELETE(":id", middleware.RoleMiddleware(models.Receptionist.String()), handler.DeletePatientByID)
	}

	http.ListenAndServe(":8080", r)
}
