package emailservice

import (
	"email-service/emailservice/models"
	"email-service/emailservice/providers"
	"errors"
	"fmt"
	"log"
)

type emailService struct {
	providers []models.EmailProvider
}

func NewEmailService() (models.EmailServiceFacade, error) {
	config := loadConfig()
	var emailProviders []models.EmailProvider
	for _, providerConfig := range config.Providers {
		provider, err := providers.CreateProvider(providerConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create provider: %v", err)
		}
		emailProviders = append(emailProviders, provider)
	}
	if len(emailProviders) == 0 {
		return nil, fmt.Errorf("no email providers configured")
	}
	return &emailService{providers: emailProviders}, nil
}

func (s *emailService) SendEmail(email *models.Email) error {
	for _, provider := range s.providers {
		err := provider.Send(email)
		if err == nil {
			return nil
		}
		log.Printf("Provider failed: %v. Attempting next provider.", err)
	}
	return errors.New("all providers failed to send the email")
}
