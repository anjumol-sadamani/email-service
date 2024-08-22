package providers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"email-service/emailservice/models"

	"github.com/mailgun/mailgun-go/v4"
)

type mailgunProvider struct {
	apiKey string
	domain string
}

func (mg *mailgunProvider) Send(email *models.Email) error {
	if mg.apiKey == "" || mg.domain == "" {
		return errors.New("mailgun api key or domain is missing")
	}

	mgClient := mailgun.NewMailgun(mg.domain, mg.apiKey)

	message := mgClient.NewMessage(
		email.From,
		email.Subject,
		email.Body,
		email.To...,
	)

	message.SetHtml(email.Body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mgClient.Send(ctx, message)

	if err != nil {
		return fmt.Errorf("failed to send email via Mailgun: %v", err)
	}

	log.Printf("Email sent via Mailgun: From %s, To %v, Subject: %s, ID: %s, Resp: %s",
		email.From, email.To, email.Subject, id, resp)
	return nil
}

func newMailgunProvider(apiKey, domain string) models.EmailProvider {
	return &mailgunProvider{apiKey: apiKey, domain: domain}
}
