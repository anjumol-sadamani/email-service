package providers

import (
	"errors"
	"fmt"
	"log"

	"email-service/emailservice/models"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridProvider struct {
	apiKey string
}

func (sg *sendGridProvider) Send(email *models.Email) error {
	if sg.apiKey == "" {
		return errors.New("SendGrid API key is missing")
	}
	from := mail.NewEmail("", email.From)
	subject := email.Subject
	to := mail.NewEmail("", email.To[0])
	plainTextContent := email.Body
	htmlContent := email.Body

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sg.apiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("failed to send email via SendGrid: %v", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("SendGrid API returned an error: %d - %s", response.StatusCode, response.Body)
	}

	log.Printf("Email sent via SendGrid: From %s, To %v, Subject: %s", email.From, email.To, email.Subject)
	return nil
}

func newSendGridProvider(apiKey string) models.EmailProvider {
	return &sendGridProvider{apiKey: apiKey}
}
