package db

import (
	"log"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/infrastructure/logger"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserEmailRepository struct {
	emailDB *sqlx.DB
}

func (r UserEmailRepository) Users() ([]domain.User, *errors.AppError) {
	var users []domain.User
	err := r.emailDB.Select(&users, "SELECT * FROM users")
	if err != nil {
		// Log the actual database error
		log.Println("Database error:", err)
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	return users, nil
}

func (r UserEmailRepository) IdNo(idNo string) (*domain.User, *errors.AppError) {
	var user domain.User
	err := r.emailDB.Get(&user, "SELECT * FROM users WHERE id_no = $1", idNo)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return &user, nil
}

func (r UserEmailRepository) CreateUser(user domain.User) (*domain.UserCreateReturn, *errors.AppError) {
	createUserSql := `
            INSERT INTO users (
                    id_no, department, first_name, last_name, suffix, email,
                    email_status, status, ticket_no, profile_picture, hashed_password,
                    salt, smtp_email, smtp_password, date_created, date_updated, created_by, updated_by
            ) VALUES (
                    :id_no, :department, :first_name, :last_name, :suffix, :email,
                    :email_status, :status, :ticket_no, :profile_picture, :hashed_password,
                    :salt, :smtp_email, :smtp_password,
                    NOW(), NOW(), :created_by, :updated_by
            )
			RETURNING id_no, first_name, last_name, email
    `

	// Use NamedQuery for named parameters with structs
	rows, err := r.emailDB.NamedQuery(createUserSql, user)
	if err != nil {
		logger.Error("Error while creating user: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()

	var userReturn domain.UserCreateReturn
	if rows.Next() {
		err = rows.StructScan(&userReturn)
		if err != nil {
			logger.Error("Error scanning user return: " + err.Error())
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after insert")
		return nil, errors.NewUnExpectedError("User creation failed")
	}

	return &userReturn, nil
}

func NewUserRepositoryDb(db *sqlx.DB) UserEmailRepository {
	return UserEmailRepository{db}
}
