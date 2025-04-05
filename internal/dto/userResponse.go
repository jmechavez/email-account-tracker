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
	TicketNo       string `json:"ticket_no,omitempty"`
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
	IdNo            string `json:"id_no"`
	Department      string `json:"department"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Suffix          string `json:"suffix"`
	Email           string `json:"email"`
	EmailStatus     string `json:"email_status"`
	Status          string `json:"status"`
	TicketNo        string `json:"ticket_no"`
	UpdatedTicketNo string `json:"updated_ticket_no"`
	DeletedTicketNo string `json:"deleted_ticket_no"`
	ProfilePicture  string `json:"profile_picture,omitempty"`
	HashedPassword  string `json:"hashed_password,omitempty"`
	Salt            string `json:"salt,omitempty"`
	SMTPEmail       string `json:"smtp_email,omitempty"`
	SMTPPassword    string `json:"smtp_password,omitempty"`
	DateCreated     string `json:"date_created"`
	DateUpdated     string `json:"date_updated"`
	DateDeleted     string `json:"date_deleted"`
	CreatedBy       string `json:"created_by"`
	UpdatedBy       string `json:"updated_by"`
	DeletedBy       string `json:"deleted_by"`
}

type UserCreateResponse struct {
	IdNo        string `json:"id_no"`
	Department  string `json:"department"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Suffix      string `json:"suffix"`
	Email       string `json:"email"`
	EmailStatus string `json:"email_status"`
	Status      string `json:"status"`
	TicketNo    string `json:"ticket_no"`
	DateCreated string `json:"date_created"`
	CreatedBy   string `json:"created_by"`
}

type UserEmailDeleteResponse struct {
	IdNo        string `json:"id_no"`
	EmailStatus string `json:"email_status"`
	Status      string `json:"status"`
}

type UserUpdateResponse struct {
	IdNo            string `json:"id_no"`
	Department      string `json:"department"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Suffix          string `json:"suffix,omitempty"`
	Email           string `json:"email"`
	EmailStatus     string `json:"email_status"`
	Status          string `json:"status"`
	UpdatedTicketNo string `json:"updated_ticket_no,omitempty"`
	ProfilePicture  string `json:"profile_picture,omitempty"`
	DateUpdated     string `json:"date_updated"`
	UpdatedBy       string `json:"updated_by"`
}

type UserUpdateSurnameResponse struct {
	IdNo            string `json:"id_no"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Suffix          string `json:"suffix,omitempty"`
	Email           string `json:"email"`
}

type UserPassCreateResponse struct {
	IdNo		   string `json:"id_no"`
	HashedPassword string `json:"hashed_password"`
	Salt           string `json:"salt"`
}
