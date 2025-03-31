package domain

import (
	"database/sql"

	"github.com/jmechavez/email-account-tracker/errors"
	"github.com/jmechavez/email-account-tracker/internal/dto"
)

type NullString struct {
	String string
	Valid  bool
}

type User struct {
	IdNo            string         `json:"id_no" db:"id_no"`
	Department      string         `json:"department" db:"department"`
	FirstName       string         `json:"first_name" db:"first_name"`
	LastName        string         `json:"last_name" db:"last_name"`
	Suffix          string         `json:"suffix" db:"suffix"`
	Email           string         `json:"email" db:"email"`
	EmailStatus     string         `json:"email_status" db:"email_status"`
	Status          string         `json:"status" db:"status"`
	TicketNo        sql.NullString `json:"ticket_no" db:"ticket_no"`
	UpdatedTicketNo sql.NullString `json:"updated_ticket_no" db:"updated_ticket_no"`
	DeletedTicketNo sql.NullString `json:"deleted_ticket_no" db:"deleted_ticket_no"`
	ProfilePicture  string         `json:"profile_picture" db:"profile_picture"`
	HashedPassword  string         `json:"hashed_password" db:"hashed_password"`
	Salt            string         `json:"salt" db:"salt"`
	SMTPEmail       string         `json:"smtp_email" db:"smtp_email"`
	SMTPPassword    string         `json:"smtp_password" db:"smtp_password"`
	DateCreated     sql.NullString `json:"date_created" db:"date_created"`
	DateUpdated     sql.NullString `json:"date_updated" db:"date_updated"`
	DateDeleted     sql.NullString `json:"date_deleted" db:"date_deleted"`
	CreatedBy       string         `json:"created_by" db:"created_by"`
	UpdatedBy       string         `json:"updated_by" db:"updated_by"`
	DeletedBy       sql.NullString `json:"deleted_by" db:"deleted_by"`
}

type UserCreateReturn struct {
	IdNo      string `json:"id_no" db:"id_no"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Suffix    string `json:"suffix" db:"suffix"`
	Email     string `json:"email" db:"email"`
}

type UserDeleteReturn struct {
	IdNo        string `json:"id_no" db:"id_no"`
	EmailStatus string `json:"email_status" db:"email_status"`
	Status      string `json:"status" db:"status"`
}

func (u User) ToDto() dto.UserEmailResponse {
	return dto.UserEmailResponse{
		IdNo:       u.IdNo,
		Department: u.Department,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Suffix:     u.Suffix,
		Email:      u.Email,
		Status:     u.Status,
	}
}

func (u User) ToIdDto() dto.UserIdNoEmailResponse {
	return dto.UserIdNoEmailResponse{
		IdNo:            u.IdNo,
		Department:      u.Department,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Suffix:          u.Suffix,
		Email:           u.Email,
		EmailStatus:     u.EmailStatus,
		Status:          u.Status,
		TicketNo:        u.TicketNo.String,
		UpdatedTicketNo: u.UpdatedTicketNo.String,
		DeletedTicketNo: u.DeletedTicketNo.String,
		ProfilePicture:  u.ProfilePicture,
		HashedPassword:  u.HashedPassword,
		Salt:            u.Salt,
		SMTPEmail:       u.SMTPEmail,
		SMTPPassword:    u.SMTPPassword,
		DateCreated:     u.DateCreated.String,
		DateUpdated:     u.DateUpdated.String,
		DateDeleted:     u.DateDeleted.String,
		CreatedBy:       u.CreatedBy,
		UpdatedBy:       u.UpdatedBy,
		DeletedBy:       u.DeletedBy.String,
	}
}

func (u User) ToNewUserDto() dto.UserCreateResponse {
	return dto.UserCreateResponse{
		IdNo:        u.IdNo,
		Department:  u.Department,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		EmailStatus: u.EmailStatus,
		Status:      u.Status,
		TicketNo:    u.TicketNo.String,
		DateCreated: u.DateCreated.String,
		CreatedBy:   u.CreatedBy,
	}
}

func (u User) ToUpdateDto() dto.UserUpdateResponse {
	return dto.UserUpdateResponse{
		IdNo:            u.IdNo,
		Department:      u.Department,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Suffix:          u.Suffix,
		Email:           u.Email,
		EmailStatus:     u.EmailStatus,
		Status:          u.Status,
		UpdatedTicketNo: u.UpdatedTicketNo.String,
		ProfilePicture:  u.ProfilePicture,
		DateUpdated:     u.DateUpdated.String,
		UpdatedBy:       u.UpdatedBy,
	}
}

func (u User) ToUserCreateReturn() UserCreateReturn {
	return UserCreateReturn{
		IdNo:      u.IdNo,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

type UserRepository interface {
	Users() ([]User, *errors.AppError)
	IdNo(string) (*User, *errors.AppError)
	CreateUser(User) (*UserCreateReturn, *errors.AppError)
	DeleteUser(User) (*UserDeleteReturn, *errors.AppError)
	UpdateUser(User) (*User, *errors.AppError)
}
