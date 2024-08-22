package emailservice

import (
	"os"

	"email-service/emailservice/models"
	"email-service/emailservice/providers"
)

func loadConfig() *models.ServiceConfig {

	return &models.ServiceConfig{
		Providers: []models.ProviderConfig{
			{
				Name: providers.SendGridProvider,
				Config: map[string]string{
					"api_key": os.Getenv("SENDGRID_API_KEY"),
				},
			},
			{
				Name: providers.MailgunProvider,
				Config: map[string]string{
					"api_key": os.Getenv("MAILGUN_API_KEY"),
					"domain":  os.Getenv("MAILGUN_DOMAIN"),
				},
			},
		},
	}
}
