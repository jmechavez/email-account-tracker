package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmechavez/email-account-tracker/internal/dto"
)

type UserService interface {
	UserNoDto() ([]domain.User, *errors.AppError)
	Users() ([]dto.UserEmailResponse, *errors.AppError)
	IdNo(idNo string) (*dto.UserIdNoEmailResponse, *errors.AppError)
	CreateUser(user dto.UserEmailRequest) (*dto.UserCreateResponse, *errors.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

// NoDto is used to return the User struct without the dto
func (s DefaultUserService) UserNoDto() ([]domain.User, *errors.AppError) {
	u, err := s.repo.Users()
	if err != nil {
		return nil, err
	}
	return u, nil
}

// User is used to return the User struct with the dto
func (s DefaultUserService) Users() ([]dto.UserEmailResponse, *errors.AppError) {
	u, err := s.repo.Users()
	if err != nil {
		return nil, err
	}
	var users []dto.UserEmailResponse
	for _, user := range u {
		users = append(users, user.ToDto())
	}
	return users, nil
}

func (s DefaultUserService) IdNo(idNo string) (*dto.UserIdNoEmailResponse, *errors.AppError) {
	u, err := s.repo.IdNo(idNo)
	if err != nil {
		return nil, err
	}

	response := u.ToIdDto()
	return &response, nil
}

func (s DefaultUserService) CreateUser(req dto.UserEmailRequest) (*dto.UserCreateResponse, *errors.AppError) {
	email, err := s.generateEmail(req.FirstName, req.LastName, req.Suffix)
	if err != nil {
		return nil, err // Assuming generateEmail returns *errors.AppError, adjust if needed.
	}

	user := domain.User{
		IdNo:           req.IdNo,
		Department:     req.Department,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Suffix:         req.Suffix,
		Email:          email,
		EmailStatus:    "active",
		Status:         req.Status,
		TicketNo:       req.TicketNo,
		ProfilePicture: "n/a",
		DateCreated:    time.Now().Format("2006-01-02 15:04:05"),
		DateUpdated:    time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:      "admin",
		UpdatedBy:      "admin",
	}

	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Create a response using the data from newUser and original user
	response := dto.UserCreateResponse{
		IdNo:        newUser.IdNo,
		Department:  user.Department,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		EmailStatus: user.EmailStatus,
		Status:      user.Status,
		TicketNo:    user.TicketNo,
		DateCreated: user.DateCreated,
		CreatedBy:   user.CreatedBy,
	}

	return &response, nil
}

func (s DefaultUserService) generateEmail(firstName, lastName, suffix string) (string, *errors.AppError) {
	// 1. Normalize the names (lowercase, remove spaces, remove dots)
	normalizedFirstName := strings.ToLower(strings.ReplaceAll(firstName, " ", ""))
	normalizedLastName := strings.ToLower(strings.ReplaceAll(lastName, " ", ""))

	// 2. Handle the suffix (if any)
	suffixPart := ""
	if suffix != "" {
		suffixPart = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(suffix, " ", ""), ".", ""))
	}

	// 3. Construct the email address
	email := fmt.Sprintf("%s.%s%s@test.com", normalizedFirstName, normalizedLastName, suffixPart)

	// 4. Validate the email (optional, but recommended)
	// Add email validation logic here if needed.

	return email, nil
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
