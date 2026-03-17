package utils

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Mailer(subject string, username string, toEmail string, textBody string, htmlBody string) error {

	from := mail.NewEmail("P2P Sharing", os.Getenv("MAIL_FROM"))

	to := mail.NewEmail(username, toEmail)

	email := mail.NewSingleEmail(
		from,
		subject,
		to,
		textBody,
		htmlBody,
	)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(email)
	if err != nil {
		log.Println("Sendgrid error:", err)
		return err
	}

	log.Println("Email sent:", response.StatusCode, response.Body, response.Headers)

	return nil
}