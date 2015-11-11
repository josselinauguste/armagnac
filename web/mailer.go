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
		fmt.Sprintf("%v:%v", os.Getenv("SMTP_SERVER_URL"), os.Getenv("SMTP_SERVER_PORT")),
		getAuth(),
		"alice@armagnac.io",
		[]string{recipient},
		msg)
	return err
}

func getHeaders(subject string) string {
	mimeHeader := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectHeader := fmt.Sprintf("Subject: %v\n", subject)
	return subjectHeader + mimeHeader
}

func getAuth() smtp.Auth {
	return smtp.PlainAuth(
		"",
		os.Getenv("SMTP_SERVER_LOGIN"),
		os.Getenv("SMTP_SERVER_PASSWORD"),
		os.Getenv("SMTP_SERVER_URL"),
	)
}
