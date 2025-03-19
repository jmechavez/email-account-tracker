package services

import (
	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/domain"
)

type UserService interface {
	User() ([]domain.User, *errors.AppError)
	
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) User() ([]domain.User, *errors.AppError) {
	u, err := s.repo.Users()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
