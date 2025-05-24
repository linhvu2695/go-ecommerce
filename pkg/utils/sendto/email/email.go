package email

import (
	"bytes"
	"fmt"
	"go-ecommerce/global"
	"net/smtp"
	"strings"
	"text/template"

	"go.uber.org/zap"
)

const (
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"

	TEMPLATES_EMAIL_FOLDER = "templates/email"
	TEMPLATE_OTP_AUTH      = "otp-auth"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Email struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(email Email) string {
	msg := "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=\"utf-8\";\r\n"
	msg += fmt.Sprintf("From: %s <%s>\r\n", email.From.Name, email.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(email.To, "; "))
	msg += fmt.Sprintf("Subject: %s\r\n", email.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", email.Body)

	return msg
}

func GetEmailTemplate(templateName string, data map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	filename := templateName + ".html"
	filepath := TEMPLATES_EMAIL_FOLDER + "/" + filename

	t := template.Must(template.New(filename).ParseFiles(filepath))

	err := t.Execute(htmlTemplate, data)
	if err != nil {
		global.Logger.Error("Failed to read template email", zap.Error(err))
		return "", err
	}

	return htmlTemplate.String(), nil
}

func SendEmailTemplate(to []string, from string, subject string, templateName string, data map[string]interface{}) error {
	htmlBody, err := GetEmailTemplate(templateName, data)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Failed to get HTML email from template: %s", templateName), zap.Error(err))
		return err
	}

	contentEmail := Email{
		From: EmailAddress{
			Address: from,
			Name:    "Ecommerce System",
		},
		To:      to,
		Subject: subject,
		Body:    htmlBody,
	}

	messageEmail := BuildMessage(contentEmail)
	global.Logger.Info("Email content", zap.String("email", messageEmail))

	// Send email SMTP
	err = smtp.SendMail(fmt.Sprintf("%s:%s", SMTP_HOST, SMTP_PORT), *global.SmtpAuth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	return nil
}

func SendEmailOtp(to []string, from string, otp int) error {
	return SendEmailTemplate(to, from,
		"Ecommerce System - OTP Verification",
		TEMPLATE_OTP_AUTH,
		map[string]interface{}{
			"otp": otp,
		})
}
