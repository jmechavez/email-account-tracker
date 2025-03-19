package domain

import (
	"time"

	"github.com/jmechavez/email-account-tracker/errors"
)

type User struct {
	IdNo           string    `json:"id_no" db:"id_no"`
	Department     string    `json:"department" db:"department"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Suffix         string    `json:"suffix" db:"suffix"`
	Email          string    `json:"email" db:"email"`
	EmailStatus    string    `json:"email_status" db:"email_status"`
	Status         string    `json:"status" db:"status"`
	TicketNo       string    `json:"ticket_no" db:"ticket_no"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	HashedPassword string    `json:"hashed_password" db:"hashed_password"`
	Salt           string    `json:"salt" db:"salt"`
	SMTPEmail      string    `json:"smtp_email" db:"smtp_email"`
	SMTPPassword   string    `json:"smtp_password" db:"smtp_password"`
	DateCreated    time.Time `json:"date_created" db:"date_created"`
	DateUpdated    time.Time `json:"date_updated" db:"date_updated"`
	DateDeleted    time.Time `json:"date_deleted" db:"date_deleted"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	UpdatedBy      string    `json:"updated_by" db:"updated_by"`
	DeletedBy      string    `json:"deleted_by" db:"deleted_by"`
}
type UserRepository interface {
	Users() ([]User, *errors.AppError)
}
