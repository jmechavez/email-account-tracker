package dto

type UserEmailResponse struct {
	IdNo           string `json:"id_no"`
	Department     string `json:"department"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Suffix         string `json:"suffix"`
	Email          string `json:"email"`
	EmailStatus    string `json:"email_status"`
	Status         string `json:"status"`
	TicketNo       string `json:"ticket_no"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	SMTPEmail      string `json:"smtp_email,omitempty"`
	SMTPPassword   string `json:"smtp_password,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateUpdated    string `json:"date_updated,omitempty"`
	DateDeleted    string `json:"date_deleted,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	UpdatedBy      string `json:"updated_by,omitempty"`
	DeletedBy      string `json:"deleted_by,omitempty"`
}

type UserIdNoEmailResponse struct {
	IdNo           string `json:"id_no"`
	Department     string `json:"department"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Suffix         string `json:"suffix"`
	Email          string `json:"email"`
	EmailStatus    string `json:"email_status"`
	Status         string `json:"status"`
	TicketNo       string `json:"ticket_no"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	SMTPEmail      string `json:"smtp_email,omitempty"`
	SMTPPassword   string `json:"smtp_password,omitempty"`
	DateCreated    string `json:"date_created"`
	DateUpdated    string `json:"date_updated"`
	DateDeleted    string `json:"date_deleted"`
	CreatedBy      string `json:"created_by"`
	UpdatedBy      string `json:"updated_by"`
	DeletedBy      string `json:"deleted_by"`
}

type UserCreateResponse struct {
	IdNo        string `json:"id_no"`
	Department  string `json:"department"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	EmailStatus string `json:"email_status"`
	Status      string `json:"status"`
	TicketNo    string `json:"ticket_no"`
	DateCreated string `json:"date_created"`
	CreatedBy   string `json:"created_by"`
}
