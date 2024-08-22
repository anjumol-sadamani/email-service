package models

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type EmailServiceFacade interface {
	SendEmail(email *Email) error
}

type EmailProvider interface {
	Send(email *Email) error
}
