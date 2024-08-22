package providers

import (
	"email-service/emailservice/models"
	"fmt"
)

func CreateProvider(config models.ProviderConfig) (models.EmailProvider, error) {
	switch config.Name {
	case SendGridProvider:
		return newSendGridProvider(config.Config["api_key"]), nil
	case MailgunProvider:
		return newMailgunProvider(config.Config["api_key"], config.Config["domain"]), nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", config.Name)
	}
}
