package db

import (
	"database/sql"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/infrastructure/logger"
	"github.com/jmechavez/email-account-tracker/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type UserEmailRepository struct {
	emailDB *sqlx.DB
}

func (r UserEmailRepository) Users(limit, offset int) ([]domain.User, *errors.AppError) {
	logger.Info("Fetching users from the database with pagination", zap.Int("limit", limit), zap.Int("offset", offset))

	var users []domain.User
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	err := r.emailDB.Select(&users, query, limit, offset)
	if err != nil {
		logger.Error("Database error while fetching users", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	logger.Info("Successfully fetched users", zap.Int("count", len(users)))
	return users, nil
}

func (r UserEmailRepository) IdNo(idNo string) (*domain.User, *errors.AppError) {
	logger.Info("Fetching user by ID", zap.String("id_no", idNo))
	var user domain.User
	err := r.emailDB.Get(&user, "SELECT * FROM users WHERE id_no = $1", idNo)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("User not found", zap.String("id_no", idNo))
			return nil, errors.NewNotFoundError("User not found")
		}
		logger.Error("Database error while fetching user by ID", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	logger.Info("Successfully fetched user", zap.String("id_no", idNo))
	return &user, nil
}

func (r UserEmailRepository) CreateUser(user domain.User) (*domain.UserCreateReturn, *errors.AppError) {
	logger.Info("Creating a new user", zap.String("id_no", user.IdNo))
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
            RETURNING id_no, first_name, last_name, suffix, email
    `

	rows, err := r.emailDB.NamedQuery(createUserSql, user)
	if err != nil {
		logger.Error("Error while creating user", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()

	var userReturn domain.UserCreateReturn
	if rows.Next() {
		err = rows.StructScan(&userReturn)
		if err != nil {
			logger.Error("Error scanning user return", zap.Error(err))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after insert")
		return nil, errors.NewUnExpectedError("User creation failed")
	}

	logger.Info("User created successfully", zap.String("id_no", userReturn.IdNo))
	return &userReturn, nil
}

func (r UserEmailRepository) DeleteUser(user domain.User) (*domain.UserDeleteReturn, *errors.AppError) {
	logger.Info("Deleting user", zap.String("id_no", user.IdNo))
	deleteUserSql := `
		UPDATE users
		SET
			status = 'deleted',
			email_status = 'deleted',
			deleted_ticket_no = $2,
			deleted_by = $3,
			date_deleted = CURRENT_TIMESTAMP
		WHERE id_no = $1 AND status != 'deleted'
		RETURNING id_no, status, email_status
	`
	var u domain.UserDeleteReturn
	err := r.emailDB.QueryRow(
		deleteUserSql,
		user.IdNo,
		user.DeletedTicketNo.String,
		user.DeletedBy.String,
	).Scan(&u.IdNo, &u.Status, &u.EmailStatus)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("User not found or already deleted", zap.String("id_no", user.IdNo))
			return nil, errors.NewNotFoundError("User not found or already deleted")
		}
		logger.Error("Database error during user deletion", zap.Error(err))
		return nil, errors.NewUnExpectedError("Database error during user deletion")
	}

	logger.Info("User deleted successfully", zap.String("id_no", u.IdNo))
	return &u, nil
}

func (r UserEmailRepository) UpdateUser(user domain.User) (*domain.User, *errors.AppError) {
	logger.Info("Updating user", zap.String("id_no", user.IdNo))
	updateUserSql := `
        UPDATE users
        SET
            department = CASE WHEN :department = '' THEN department ELSE :department END,
            first_name = CASE WHEN :first_name = '' THEN first_name ELSE :first_name END,
            last_name = CASE WHEN :last_name = '' THEN last_name ELSE :last_name END,
            suffix = CASE WHEN :suffix = '' THEN suffix ELSE :suffix END,
            email = CASE WHEN :email = '' THEN email ELSE :email END,
            email_status = CASE WHEN :email_status = '' THEN email_status ELSE :email_status END,
            status = CASE WHEN :status = '' THEN status ELSE :status END,
            ticket_no = CASE WHEN :ticket_no = '' THEN ticket_no ELSE :ticket_no END,
            profile_picture = CASE WHEN :profile_picture = '' THEN profile_picture ELSE :profile_picture END,
            updated_by = :updated_by,
            date_updated = CURRENT_TIMESTAMP
        WHERE id_no = :id_no
        RETURNING *
    `
	rows, err := r.emailDB.NamedQuery(updateUserSql, user)
	if err != nil {
		logger.Error("Error while updating user", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()

	var updatedUser domain.User
	if rows.Next() {
		err = rows.StructScan(&updatedUser)
		if err != nil {
			logger.Error("Error scanning updated user", zap.Error(err))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after update")
		return nil, errors.NewUnExpectedError("User update failed")
	}

	logger.Info("User updated successfully", zap.String("id_no", updatedUser.IdNo))
	return &updatedUser, nil
}

func (r UserEmailRepository) UpdateSurname(user domain.User) (*domain.User, *errors.AppError) {
	logger.Info("Updating user surname", zap.String("id_no", user.IdNo))
	updateSurnameSql := `
		UPDATE users
		SET
			last_name = :last_name,
			updated_ticket_no = :updated_ticket_no,
			email = :email,
			updated_by = :updated_by,
			date_updated = CURRENT_TIMESTAMP
		WHERE id_no = :id_no
		RETURNING *
	`
	rows, err := r.emailDB.NamedQuery(updateSurnameSql, user)
	if err != nil {
		logger.Error("Error while updating surname", zap.Error(err))
		return nil, errors.NewUnExpectedError("Unexpected database error")
	}
	defer rows.Close()

	var updatedUser domain.User
	if rows.Next() {
		err = rows.StructScan(&updatedUser)
		if err != nil {
			logger.Error("Error scanning updated user", zap.Error(err))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	} else {
		logger.Error("No rows returned after surname update")
		return nil, errors.NewUnExpectedError("Surname update failed")
	}

	logger.Info("User surname updated successfully", zap.String("id_no", updatedUser.IdNo))
	return &updatedUser, nil
}

func NewUserRepositoryDb(db *sqlx.DB) UserEmailRepository {
	logger.Info("Initializing UserEmailRepository")
	return UserEmailRepository{db}
}
