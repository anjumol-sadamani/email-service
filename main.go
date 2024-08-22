package main

import (
	"email-service/emailservice"
	"email-service/emailservice/models"
	"log"

	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	emailClient, err := emailservice.NewEmailService()
	if err != nil {
		log.Fatalf("Failed to create email service: %v", err)
	}

	from := flag.String("from", "", "Sender email address")
	to := flag.String("to", "", "Recipient email address")

	flag.Parse()

	if *from == "" || *to == "" {
		fmt.Println("All arguments (from, to) are required")
		flag.Usage()
		return
	}

	email := &models.Email{
		From:    *from,
		To:      []string{*to},
		Subject: "Test Email",
		Body:    "This is a test email.",
	}

	err = emailClient.SendEmail(email)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}
