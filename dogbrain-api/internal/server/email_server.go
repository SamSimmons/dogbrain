package server

import (
	"fmt"

	"github.com/keighl/postmark"
)

type EmailService struct {
    client *postmark.Client
    from   string
}

func NewEmailService(serverToken, accountToken, fromEmail string) *EmailService {
    client := postmark.NewClient(serverToken, accountToken)
    return &EmailService{
        client: client,
        from:   fromEmail,
    }
}

func (e *EmailService) SendVerificationEmail(toEmail, token string) error {
    email := postmark.Email{
        From:       e.from,
        To:         toEmail,
        Subject:    "Please verify your email address",
        HtmlBody:   fmt.Sprintf(`
            <h1>Welcome!</h1>
            <p>Please verify your email address by clicking the link below:</p>
            <p><a href="http://localhost:5173/verify/%s">Verify Email</a></p>
            <p>This link will expire in 24 hours.</p>
        `, token),
        TextBody:   fmt.Sprintf("Welcome! Please verify your email address by visiting this link: http://localhost:8080/verify/%s", token),
        Tag:        "verification",
        TrackOpens: false,
    }

    _, err := e.client.SendEmail(email)
    return err
}