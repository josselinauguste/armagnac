package web

import (
	"fmt"
	"net/smtp"
	"os"
)

type mailer interface {
	sendMail(recipient string, subject string, content []byte) error
}

type htmlMailer struct{}

func (htmlMailer) sendMail(recipient string, subject string, content []byte) error {
	headers := getHeaders(subject)
	msg := append([]byte(headers), content...)
	err := smtp.SendMail(
		getServerAddress(),
		getAuth(),
		os.Getenv("SENDER_EMAIL"),
		[]string{recipient},
		msg)
	return err
}

func getHeaders(subject string) string {
	mimeHeader := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectHeader := fmt.Sprintf("Subject: %v\n", subject)
	return subjectHeader + mimeHeader
}

func getServerAddress() string {
	smtpUrl := os.Getenv("SMTP_SERVER_URL")
	if smtpUrl == "" {
		smtpUrl = "localhost"
	}
	smtpPort := os.Getenv("SMTP_SERVER_PORT")
	if smtpPort == "" {
		smtpPort = "25"
	}
	return fmt.Sprintf("%v:%v", smtpUrl, smtpPort)
}

func getAuth() smtp.Auth {
	if os.Getenv("SMTP_SERVER_LOGIN") == "" {
		return nil
	}
	return smtp.PlainAuth(
		"",
		os.Getenv("SMTP_SERVER_LOGIN"),
		os.Getenv("SMTP_SERVER_PASSWORD"),
		os.Getenv("SMTP_SERVER_URL"),
	)
}
