package repository

import (
	"errors"

	"github.com/20ritiksingh/hospital-app/internal/utils"
	"github.com/20ritiksingh/hospital-app/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CreateUser(user *models.User) (*models.User, error) {
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	// Create the user in the database
	// Check if the email already exists
	existingUser, err := r.FindByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != gorm.ErrRecordNotFound {
		return existingUser, errors.New("user with this email already exists")
	}
	// Create the user
	err = r.db.Create(user).Error
	return existingUser, err
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
