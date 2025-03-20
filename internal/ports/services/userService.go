package services

import (
	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmechavez/email-account-tracker/internal/dto"
)

type UserService interface {
	UserNoDto() ([]domain.User, *errors.AppError)
	Users() ([]dto.UserEmailResponse, *errors.AppError)
	IdNo(idNo string) (*dto.UserIdNoEmailResponse, *errors.AppError)
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

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
