package core

// EmailSender is the gateway to send emails
type EmailSender interface {
	SendEmail(to string, subject string, body string)
}
