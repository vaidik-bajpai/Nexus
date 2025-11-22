package mailer

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/smtp"
	"time"

	"github.com/vaidik-bajpai/Nexus/backend/internal/helper"
)

//go:embed templates/verify-email.tmpl
var verifyEmailTemplate string

//go:embed templates/password-reset.tmpl
var passwordResetTemplate string

var emailTemplate *template.Template
var passwordResetEmailTemplate *template.Template

func init() {
	var err error
	emailTemplate, err = template.New("verify-email").Parse(verifyEmailTemplate)
	if err != nil {
		panic("failed to parse email template: " + err.Error())
	}

	passwordResetEmailTemplate, err = template.New("password-reset").Parse(passwordResetTemplate)
	if err != nil {
		panic("failed to parse password reset email template: " + err.Error())
	}
}

// Mailer defines the interface for sending emails
type Mailer interface {
	SendEmailVerificationEmail(to []string, subject string, verificationURL string) error
	SendPasswordResetEmail(to []string, subject string, passwordResetURL string) error
}

// SMTPMailer implements the Mailer interface using SMTP
type SMTPMailer struct{}

// NewSMTPMailer creates a new SMTP mailer instance
func NewSMTPMailer() *SMTPMailer {
	return &SMTPMailer{}
}

type EmailData struct {
	VerificationURL string
	Year            int
}
type PasswordResetData struct {
	PasswordResetURL string
	Year             int
}

// SendEmailVerificationEmail sends an email verification email
func (m *SMTPMailer) SendEmailVerificationEmail(to []string, subject string, verificationURL string) error {
	data := EmailData{
		VerificationURL: verificationURL,
		Year:            time.Now().Year(),
	}

	var buf bytes.Buffer
	if err := emailTemplate.Execute(&buf, data); err != nil {
		return err
	}

	htmlBody := buf.String()

	fromEmail := helper.GetStrEnvOrPanic("FROM_EMAIL")

	// Format email message with proper headers
	msg := fmt.Sprintf("From: %s\r\n", fromEmail)
	msg += fmt.Sprintf("To: %s\r\n", to[0])
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=UTF-8\r\n"
	msg += "\r\n"
	msg += htmlBody

	auth := smtp.PlainAuth(
		"",
		fromEmail,
		helper.GetStrEnvOrPanic("FROM_EMAIL_PASSWORD"),
		helper.GetStrEnvOrPanic("FROM_EMAIL_SMTP"),
	)

	return smtp.SendMail(
		helper.GetStrEnvOrPanic("SMTP_ADDR"),
		auth,
		fromEmail,
		to,
		[]byte(msg),
	)
}

// SendPasswordResetEmail sends a password reset email
func (m *SMTPMailer) SendPasswordResetEmail(to []string, subject string, passwordResetURL string) error {
	data := PasswordResetData{
		PasswordResetURL: passwordResetURL,
		Year:             time.Now().Year(),
	}

	var buf bytes.Buffer
	if err := passwordResetEmailTemplate.Execute(&buf, data); err != nil {
		return err
	}

	htmlBody := buf.String()

	fromEmail := helper.GetStrEnvOrPanic("FROM_EMAIL")

	// Format email message with proper headers
	msg := fmt.Sprintf("From: %s\r\n", fromEmail)
	msg += fmt.Sprintf("To: %s\r\n", to[0])
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=UTF-8\r\n"
	msg += "\r\n"
	msg += htmlBody

	auth := smtp.PlainAuth(
		"",
		fromEmail,
		helper.GetStrEnvOrPanic("FROM_EMAIL_PASSWORD"),
		helper.GetStrEnvOrPanic("FROM_EMAIL_SMTP"),
	)

	return smtp.SendMail(
		helper.GetStrEnvOrPanic("SMTP_ADDR"),
		auth,
		fromEmail,
		to,
		[]byte(msg),
	)
}

// SendEmailVerificationEmail is a convenience function that uses the default SMTP mailer
// Deprecated: Use Mailer interface instead
func SendEmailVerificationEmail(to []string, subject string, verificationURL string) error {
	mailer := NewSMTPMailer()
	return mailer.SendEmailVerificationEmail(to, subject, verificationURL)
}

// SendPasswordResetEmail is a convenience function that uses the default SMTP mailer
// Deprecated: Use Mailer interface instead
func SendPasswordResetEmail(to []string, subject string, passwordResetURL string) error {
	mailer := NewSMTPMailer()
	return mailer.SendPasswordResetEmail(to, subject, passwordResetURL)
}
