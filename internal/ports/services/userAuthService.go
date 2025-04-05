package services

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmechavez/email-account-tracker/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthService interface {
	CreatePassword(user dto.UserPassCreateRequest) (*dto.UserPassCreateResponse, *errors.AppError)
}

type DefaultUserAuthService struct {
	repo domain.UserAuthRepository
}

func (s DefaultUserAuthService) CreatePassword(req dto.UserPassCreateRequest) (*dto.UserPassCreateResponse, *errors.AppError) {
	// Generate hashed password
	hashedPassword, err := s.GenerateHashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Generate salt
	salt, err := s.GenerateSalt()
	if err != nil {
	return nil, err
	}

	user := domain.User{
		IdNo:           req.IdNo,
		HashedPassword: hashedPassword,
		Salt:           salt,
	}

	securePassword, err := s.repo.CreatePassword(user)
	if err != nil {
		return nil, err
	}

	log.Printf("User with ID %s successfully created a password", securePassword.IdNo)

	// Create response
	response := &dto.UserPassCreateResponse{
		IdNo:           user.IdNo,
		HashedPassword: user.HashedPassword,
		Salt:           user.Salt,
	}

	return response, nil
}

func (s DefaultUserAuthService) GenerateHashedPassword(password string) (string, *errors.AppError) {

	if password == "" {
		return "", errors.NewBadRequestError("Password cannot be empty")
	}

	if len(password) < 8 {
		return "", errors.NewBadRequestError("Password must be at least 8 characters long")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.NewUnExpectedError("Error generating hashed password")
	}
	return string(bytes), nil
}

// GenerateSalt generates a random salt
func (s DefaultUserAuthService) GenerateSalt() (string, *errors.AppError) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.NewUnExpectedError("Error generating salt")
	}
	// Convert to base64 string for safe storage
	return base64.StdEncoding.EncodeToString(salt), nil
}

func NewUserAuthService(repository domain.UserAuthRepository) DefaultUserAuthService {
	return DefaultUserAuthService{repository}
}
