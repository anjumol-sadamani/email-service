package providers

import (
	"email-service/emailservice/models"
	"testing"
)

func TestNewSendGridProvider(t *testing.T) {
	apiKey := "test_api_key"
	provider := newSendGridProvider(apiKey)

	sgProvider, ok := provider.(*sendGridProvider)
	if !ok {
		t.Fatal("Expected *sendGridProvider type")
	}

	if sgProvider.apiKey != apiKey {
		t.Errorf("Expected API key %s, got %s", apiKey, sgProvider.apiKey)
	}

}

func TestSendGridProvider_Send_MissingAPIKey(t *testing.T) {
	sg := &sendGridProvider{
		apiKey: "",
	}
	email := &models.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	err := sg.Send(email)
	if err == nil || err.Error() != "SendGrid API key is missing" {
		t.Fatalf("expected error 'SendGrid API key is missing', got %v", err)
	}
}
