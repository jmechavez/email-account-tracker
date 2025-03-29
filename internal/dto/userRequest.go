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
