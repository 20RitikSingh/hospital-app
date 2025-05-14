package service

import (
	"errors"

	"github.com/20ritiksingh/hospital-app/internal/config"
	"github.com/20ritiksingh/hospital-app/internal/repository"
	"github.com/20ritiksingh/hospital-app/internal/utils"
	"github.com/20ritiksingh/hospital-app/models"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// Register a new user
func (s *AuthService) Register(user models.User) (string, error) {
	// Save user to the database through the repository
	generatedUser, err := s.repo.CreateUser(&user)
	if err != nil {
		return "", err
	}
	// Generate JWT token
	token, err := utils.GenerateToken(generatedUser.IDString(), generatedUser.Role.String(), config.GetJWTSecret(), config.GetJWTExpiration())
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login user
func (s *AuthService) Login(u models.User) (string, error) {
	email := u.Email
	password := u.Password
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check password hash
	vaild, err := utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("incorrect password")
	}
	if !vaild {
		return "", errors.New("incorrect password")
	}
	// Generate JWT token
	token, err := utils.GenerateToken(user.IDString(), user.Role.String(), config.GetJWTSecret(), config.GetJWTExpiration())
	return token, err
}
