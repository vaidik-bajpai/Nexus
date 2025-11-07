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

type EmailData struct {
	VerificationURL string
	Year            int
}
type PasswordResetData struct {
	PasswordResetURL string
	Year             int
}

func SendEmailVerificationEmail(to []string, subject string, verificationURL string) error {
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

func SendPasswordResetEmail(to []string, subject string, passwordResetURL string) error {
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
