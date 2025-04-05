package dto

type UserEmailRequest struct {
	IdNo           string `json:"id_no" db:"id_no"`
	Department     string `json:"department" db:"department"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Suffix         string `json:"suffix" db:"suffix"`
	Email          string `json:"email" db:"email"`
	EmailStatus    string `json:"email_status" db:"email_status"`
	Status         string `json:"status" db:"status"`
	TicketNo       string `json:"ticket_no" db:"ticket_no"`
	ProfilePicture string `json:"profile_picture" db:"profile_picture"`
	CreatedBy      string `json:"created_by" db:"created_by"`
}

type UserEmailDeleteRequest struct {
	IdNo            string `json:"id_no" db:"id_no"`
	DeletedTicketNo string `json:"deleted_ticket_no" db:"deleted_ticket_no"`
	DeletedBy       string `json:"deleted_by" db:"deleted_by"`
}

type UserUpdateSurnameRequest struct {
	IdNo            string `json:"id_no" db:"id_no"`
	FirstName       string `json:"first_name" db:"first_name"`
	LastName        string `json:"last_name" db:"last_name"`
	Suffix          string `json:"suffix" db:"suffix"`
	UpdatedTicketNo string `json:"updated_ticket_no" db:"updated_ticket_no"`
	UpdatedBy       string `json:"updated_by" db:"updated_by"`
}

type UserUpdateRequest struct {
	IdNo            string `json:"id_no" db:"id_no"`
	Department      string `json:"department" db:"department"`
	FirstName       string `json:"first_name" db:"first_name"`
	LastName        string `json:"last_name" db:"last_name"`
	Suffix          string `json:"suffix" db:"suffix"`
	Email           string `json:"email" db:"email"`
	EmailStatus     string `json:"email_status" db:"email_status"`
	Status          string `json:"status" db:"status"`
	UpdatedTicketNo string `json:"updated_ticket_no" db:"updated_ticket_no"`
	ProfilePicture  string `json:"profile_picture" db:"profile_picture"`
	UpdatedBy       string `json:"updated_by" db:"updated_by"`
}

type UserPassCreateRequest struct {
	IdNo           string `json:"id_no" db:"id_no"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
	Salt           string `json:"salt" db:"salt"`
}

// func (r UserEmailRequest) Validate() *errors.AppError {
// 	if r.EmailAction != "save" {
// 		return errors.NewValidationError("User action failed")
// 	}
// 	return nil
// }
