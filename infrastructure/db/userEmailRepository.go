package db

import (
	"log"

	"github.com/jmechavez/email-account-tracker/errors"
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

func NewUserRepositoryDb(db *sqlx.DB) UserEmailRepository {
	return UserEmailRepository{db}
}
