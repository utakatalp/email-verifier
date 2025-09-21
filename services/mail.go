package services

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Gmail SMTP örneği (App Password gerekiyor)
/*var (
	smtpUser    = os.Getenv("SMTP_USER")
	smtpPass    = os.Getenv("SMTP_PASS")
	smtpHost    = os.Getenv("SMTP_HOST")
	smtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
)
*/
func SendActivationMail(to, token string) error {
	from := mail.NewEmail("Email Verifier", os.Getenv("FROM_EMAIL"))
	subject := "Activate your account"
	toEmail := mail.NewEmail("", to)

	baseURL := os.Getenv("PUBLIC_DOMAIN")

	link := fmt.Sprintf("%s/activate?token=%s", baseURL, token)
	plainTextContent := "Click this link to activate your account: " + link
	htmlContent := fmt.Sprintf("<p>Click <a href=\"%s\">here</a> to activate your account.</p>", link)

	message := mail.NewSingleEmail(from, subject, toEmail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	client.Request, _ = sendgrid.SetDataResidency(client.Request, "eu")
	fmt.Println(message)

	resp, err := client.Send(message)
	if err != nil {
		return ErrMailSendFailed
	}
	fmt.Printf("SendGrid Response: %d %s\n", resp.StatusCode, resp.Body)
	return nil
}
