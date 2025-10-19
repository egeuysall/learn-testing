package services

import (
	"github.com/resend/resend-go/v2"
)

type EmailSender interface {
	SendWelcomeEmail(to, name string) error
}

type ResendEmailSender struct {
	client *resend.Client
	from   string
}

func NewResendEmailSender(apiKey, fromEmail string) *ResendEmailSender {
	client := resend.NewClient(apiKey)
	return &ResendEmailSender{
		client: client,
		from:   fromEmail,
	}
}

func (r *ResendEmailSender) SendWelcomeEmail(to, name string) error {
	params := &resend.SendEmailRequest{
		From:    r.from,
		To:      []string{to},
		Subject: "Welcome to our SaaS!",
		Html:    "<h1>Welcome " + name + "!</h1><p>Thanks for signing up.</p>",
	}

	_, err := r.client.Emails.Send(params)
	return err
}
