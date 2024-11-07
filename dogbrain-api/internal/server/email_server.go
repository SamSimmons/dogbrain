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
		From:    e.from,
		To:      toEmail,
		Subject: "Please verify your email address",
		HtmlBody: fmt.Sprintf(`
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

func (e *EmailService) SendPasswordResetEmail(toEmail, token string) error {
	email := postmark.Email{
		From:    e.from,
		To:      toEmail,
		Subject: "Reset your password",
		HtmlBody: fmt.Sprintf(`
            <h1>Password Reset Request</h1>
            <p>A password reset has been requested for your account. Click the link below to reset your password:</p>
            <p><a href="http://localhost:5173/reset-password/%s">Reset Password</a></p>
            <p>This link will expire in 1 hour.</p>
            <p>If you didn't request this password reset, you can safely ignore this email.</p>
        `, token),
		TextBody:   fmt.Sprintf("A password reset has been requested for your account. Visit this link to reset your password: http://localhost:5173/reset-password/%s\n\nThis link will expire in 1 hour. If you didn't request this reset, you can safely ignore this email.", token),
		Tag:        "password-reset",
		TrackOpens: false,
	}

	_, err := e.client.SendEmail(email)
	return err
}
