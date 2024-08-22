package providers

import (
	"context"
	"email-service/emailservice/models"
	"testing"

	"github.com/mailgun/mailgun-go/v4"
)

type MockMailgunClient struct {
	sendFunc func(ctx context.Context, message *mailgun.Message) (string, string, error)
}

func (m *MockMailgunClient) Send(ctx context.Context, message *mailgun.Message) (string, string, error) {
	return m.sendFunc(ctx, message)
}

func TestNewMailgunProvider(t *testing.T) {
	apiKey := "test_api_key"
	domain := "test.com"
	provider := newMailgunProvider(apiKey, domain)

	mgProvider, ok := provider.(*mailgunProvider)
	if !ok {
		t.Fatal("Expected *mailgunProvider type")
	}

	if mgProvider.apiKey != apiKey {
		t.Errorf("Expected API key %s, got %s", apiKey, mgProvider.apiKey)
	}

	if mgProvider.domain != domain {
		t.Errorf("Expected domain %s, got %s", domain, mgProvider.domain)
	}
}

func TestMailgunProviderSendNoAPIKey(t *testing.T) {
	provider := &mailgunProvider{
		apiKey: "",
		domain: "test.com",
	}

	email := &models.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Body:    "This is a test email.",
	}

	err := provider.Send(email)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if err.Error() != "mailgun api key or domain is missing" {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}

func TestMailgunProviderSendNoDomain(t *testing.T) {
	provider := &mailgunProvider{
		apiKey: "test_api_key",
		domain: "",
	}

	email := &models.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Body:    "This is a test email.",
	}

	err := provider.Send(email)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if err.Error() != "mailgun api key or domain is missing" {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}
