package db

import (
	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/infrastructure/logger"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type UserAuthRepository struct {
	emailDB *sqlx.DB
}

func (r UserAuthRepository) Login(user domain.User) (*domain.User, *errors.AppError) {
	loginSql := `
		SELECT * FROM users
		WHERE id_no = :id_no
		AND hashed_password = :hashed_password
		AND salt = :salt
	`
	rows, err := r.emailDB.NamedQuery(loginSql, user)
	if err != nil {
		logger.Error("Error while logging in", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()
	var loggedInUser domain.User
	if rows.Next() {
		err = rows.StructScan(&loggedInUser)
		if err != nil {
			logger.Error("Error scanning logged-in user", zap.Error(err))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after login")
		return nil, errors.NewBadRequestError("Invalid login credentials")
	}
	logger.Info("User login success", zap.String("id_no", loggedInUser.IdNo))
	return &loggedInUser, nil
}

func (r UserAuthRepository) CreatePassword(user domain.User) (*domain.User, *errors.AppError) {
	passwordUserSql := `
		UPDATE users
		SET
			hashed_password = :hashed_password,
			salt = :salt
		WHERE id_no = :id_no
		RETURNING *
	`
	rows, err := r.emailDB.NamedQuery(passwordUserSql, user)
	if err != nil {
		logger.Error("Error while creating password", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()

	var updatedPassword domain.User
	if rows.Next() {
		err = rows.StructScan(&updatedPassword)
		if err != nil {
			logger.Error("Error scanning updated user", zap.Error(err))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after password update")
		return nil, errors.NewUnExpectedError("Password creation failed")
	}

	logger.Info("User password success", zap.String("id_no", updatedPassword.IdNo))
	return &updatedPassword, nil
}

func NewUserAuthRepositoryDb(db *sqlx.DB) UserAuthRepository {
	logger.Info("Initializing UserAuthRepository")
	return UserAuthRepository{db}
}
