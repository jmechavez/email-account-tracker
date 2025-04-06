package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmechavez/email-account-tracker/internal/dto"
)

// UserAuthService defines the interface for user authentication services
type UserAuthService interface {
	// CreatePassword creates a hashed password for a user
	CreatePassword(user dto.UserPassCreateRequest) (*dto.UserPassCreateResponse, *errors.AppError)
	Login(user dto.UserPassLoginRequest) (*dto.UserPassLoginResponse, *errors.AppError)
}

// DefaultUserAuthService is the default implementation of UserAuthService
type DefaultUserAuthService struct {
	repo  domain.UserAuthRepository // Repository for user authentication
	urepo domain.UserRepository     // Repository for user data
}

// Validate the password
func (s DefaultUserAuthService) Login(req dto.UserPassLoginRequest) (*dto.UserPassLoginResponse, *errors.AppError) {
	// You're missing the code to fetch the user
	existingUser, err := s.urepo.IdNo(req.IdNo)
	if err != nil {
		return nil, errors.NewUnExpectedError("Error fetching user")
	}

	if req.Password == "" {
		return nil, errors.NewBadRequestError("Password cannot be empty")
	}
	if len(req.Password) < 8 {
		return nil, errors.NewBadRequestError("Password must be at least 8 characters long")
	}

	// Hash the provided password with the stored salt
	hashedPassword, err := s.GenerateHashedPassword(req.Password, existingUser.Salt)
	if err != nil {
		return nil, errors.NewUnExpectedError("Error hashing password")
	}

	// Compare the hashes
	if hashedPassword != existingUser.HashedPassword {
		return nil, errors.NewAuthorizationError("Invalid credentials")
	}

	log.Printf("User with ID %s successfully logged in", req.IdNo)

	response := &dto.UserPassLoginResponse{
		IdNo:      existingUser.IdNo,
		FirstName: existingUser.FirstName,
	}

	return response, nil
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

	// Generate a random salt
	salt, err := s.GenerateSalt()
	if err != nil {
		return nil, err
	}

	// Generate a hashed password with the salt
	hashedPassword, err := s.GenerateHashedPassword(req.Password, salt)
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

// GenerateSalt generates a random salt
func (s DefaultUserAuthService) GenerateSalt() (string, *errors.AppError) {
	salt := make([]byte, 32) // 32 bytes salt
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.NewUnExpectedError("Error generating salt")
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// GenerateHashedPassword generates a hashed password using the provided salt
func (s DefaultUserAuthService) GenerateHashedPassword(password string, salt string) (string, *errors.AppError) {
	// Combine password with salt
	saltedPassword := password + salt

	// Use a strong hashing algorithm (here using SHA-512)
	hash := sha512.New()
	_, err := hash.Write([]byte(saltedPassword))
	if err != nil {
		return "", errors.NewUnExpectedError("Error hashing password")
	}

	hashedBytes := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashedBytes), nil
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
