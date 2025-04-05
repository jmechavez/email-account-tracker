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

// UserAuthService defines the interface for user authentication services
type UserAuthService interface {
	// CreatePassword creates a hashed password for a user
	CreatePassword(user dto.UserPassCreateRequest) (*dto.UserPassCreateResponse, *errors.AppError)
}

// DefaultUserAuthService is the default implementation of UserAuthService
type DefaultUserAuthService struct {
	repo  domain.UserAuthRepository // Repository for user authentication
	urepo domain.UserRepository     // Repository for user data
}

// CreatePassword creates a hashed password for a user
func (s DefaultUserAuthService) CreatePassword(req dto.UserPassCreateRequest) (*dto.UserPassCreateResponse, *errors.AppError) {

	// Fetch the user by ID number
	existingUser, err := s.urepo.IdNo(req.IdNo)
	if err != nil {
		return nil, errors.NewUnExpectedError("Error fetching user")
	}

	// Check if the user already has a password
	if existingUser.HashedPassword != "" && existingUser.Salt != "" {
		return nil, errors.NewBadRequestError("User already has a password")
	}

	// Validate the password
	if req.Password == "" {
		return nil, errors.NewBadRequestError("Password cannot be empty")
	}
	if len(req.Password) < 8 {
		return nil, errors.NewBadRequestError("Password must be at least 8 characters long")
	}

	// Generate a hashed password
	hashedPassword, err := s.GenerateHashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Generate a random salt
	salt, err := s.GenerateSalt()
	if err != nil {
		return nil, err
	}

	// Create a user object with the hashed password and salt
	user := domain.User{
		IdNo:           req.IdNo,
		HashedPassword: hashedPassword,
		Salt:           salt,
	}

	// Save the secure password in the repository
	securePassword, err := s.repo.CreatePassword(user)
	if err != nil {
		return nil, err
	}

	// Log the successful password creation
	log.Printf("User with ID %s successfully created a password", securePassword.IdNo)

	// Create a response object
	response := &dto.UserPassCreateResponse{
		IdNo:           user.IdNo,
		HashedPassword: user.HashedPassword,
		Salt:           user.Salt,
	}

	return response, nil
}

// GenerateHashedPassword generates a hashed password using bcrypt
func (s DefaultUserAuthService) GenerateHashedPassword(password string) (string, *errors.AppError) {

	// Validate the password
	if password == "" {
		return "", errors.NewBadRequestError("Password cannot be empty")
	}
	if len(password) < 8 {
		return "", errors.NewBadRequestError("Password must be at least 8 characters long")
	}

	// Generate the hashed password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.NewUnExpectedError("Error generating hashed password")
	}
	return string(bytes), nil
}

// GenerateSalt generates a random salt for password hashing
func (s DefaultUserAuthService) GenerateSalt() (string, *errors.AppError) {
	// Create a byte slice for the salt
	salt := make([]byte, 16)

	// Generate random bytes for the salt
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.NewUnExpectedError("Error generating salt")
	}

	// Convert the salt to a base64-encoded string for safe storage
	return base64.StdEncoding.EncodeToString(salt), nil
}

// NewUserAuthService creates a new instance of DefaultUserAuthService
func NewUserAuthService(
	authRepo domain.UserAuthRepository,
	userRepo domain.UserRepository,
) DefaultUserAuthService {
	return DefaultUserAuthService{
		repo:  authRepo,
		urepo: userRepo,
	}
}
