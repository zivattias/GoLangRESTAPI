package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/resendlabs/resend-go"
)

func sendEmail(c *gin.Context) (resend.SendEmailResponse, error) {
	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Ziv <hi@zivattias.com>",
		To:      []string{"ziv.attias7@gmail.com"},
		Html:    fmt.Sprintf("List of albums:\n %+v", albums),
		Subject: "Hello from Golang!",
	}

	sent, err := client.Emails.Send(params)

	if err != nil {
		return resend.SendEmailResponse{}, err
	}
	return sent, nil
}
