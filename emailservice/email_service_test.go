package emailservice

import (
	"email-service/emailservice/models"
	"errors"
	"os"
	"testing"
)

type MockProvider struct {
	sendFunc func(*models.Email) error
}

func (m *MockProvider) Send(email *models.Email) error {
	return m.sendFunc(email)
}

func TestNewEmailService(t *testing.T) {
	os.Setenv("SENDGRID_API_KEY", "test_sendgrid_key")
	os.Setenv("MAILGUN_API_KEY", "test_mailgun_key")
	os.Setenv("MAILGUN_DOMAIN", "test.com")

	service, err := NewEmailService()
	if err != nil {
		t.Fatalf("Failed to create email service: %v", err)
	}

	if service == nil {
		t.Fatal("Email service is nil")
	}
}

func TestSendEmail(t *testing.T) {
	mockProvider := &MockProvider{
		sendFunc: func(email *models.Email) error {
			if email.From != "sender@example.com" {
				t.Errorf("Expected From to be 'sender@example.com', got '%s'", email.From)
			}
			if email.To[0] != "recipient@example.com" {
				t.Errorf("Expected To to be 'recipient@example.com', got '%s'", email.To[0])
			}
			if email.Subject != "Test Email" {
				t.Errorf("Expected Subject to be 'Test Email', got '%s'", email.Subject)
			}
			if email.Body != "This is a test email." {
				t.Errorf("Expected Body to be 'This is a test email.', got '%s'", email.Body)
			}
			return nil
		},
	}

	service := &emailService{
		providers: []models.EmailProvider{mockProvider},
	}

	email := &models.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Body:    "This is a test email.",
	}

	err := service.SendEmail(email)
	if err != nil {
		t.Fatalf("Failed to send email: %v", err)
	}
}

func TestSendEmailFailure(t *testing.T) {
	mockProvider := &MockProvider{
		sendFunc: func(email *models.Email) error {
			return errors.New("mock provider error")
		},
	}

	service := &emailService{
		providers: []models.EmailProvider{mockProvider},
	}

	email := &models.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Body:    "This is a test email.",
	}

	err := service.SendEmail(email)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expectedError := "all providers failed to send the email"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}
