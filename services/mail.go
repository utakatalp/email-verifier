package services

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Gmail SMTP örneği (App Password gerekiyor)
var (
	smtpUser    = os.Getenv("SMTP_USER")
	smtpPass    = os.Getenv("SMTP_PASS")
	smtpHost    = os.Getenv("SMTP_HOST")
	smtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
)

func SendActivationMail(to string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Activate your account")

	link := fmt.Sprintf("http://localhost:8080/activate?token=%s", token)
	m.SetBody("text/plain", "Click the link to activate your email: "+link)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	return d.DialAndSend(m)
}
