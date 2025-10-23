package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	// DefaultEmailSender is the default email address used to send emails
	DefaultEmailSender = "support@lenslocked.com"
)

// EmailService is used for the email service
type EmailService struct {
	DefaultEmailSender string
	dialer             *mail.Dialer
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Email struct {
	From      string
	To        string
	Subject   string
	Plaintext string
	HTML      string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(
			config.Host,
			config.Port,
			config.Username,
			config.Password),
	}

	return &es
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	es.setFrom(msg, email)
	msg.SetHeader("Subject", email.Subject)

	// display plain text or html depending on what is available
	switch {
	case email.Plaintext != "" && email.HTML != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	case email.HTML != "":
		msg.AddAlternative("text/html", email.HTML)
	}

	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send message err: %w", err)
	}

	return nil
}

func (es *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case es.DefaultEmailSender != "":
		from = es.DefaultEmailSender
	default:
		from = DefaultEmailSender
	}

	msg.SetHeader("From", from)
}

func (es *EmailService) ForgotPassword(to, resetURL string) error {
	email := Email{
		Subject:   "Reset your password",
		To:        to,
		Plaintext: "To reset your password, please click the following link: " + resetURL,
		HTML:      `<p>To reset your password, please click the following link: <a href="` + resetURL + `">` + resetURL + `</a></p>`,
	}
	err := es.Send(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}

	return nil
}
