package subscription

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/rickferrdev/sketch-service-email/pkg/mail"
)

var (
	ErrEmptyEmail   = errors.New("email address cannot be empty")
	ErrFailedToSend = errors.New("failed to send subscription email")
)

type Service struct {
	mail *mail.Mail
}

func New(mail *mail.Mail) *Service {
	return &Service{
		mail: mail,
	}
}

func (se *Service) Signature(ctx context.Context, email string) error {
	if email == "" {
		slog.Warn("signature attempt with empty email", slog.String("context", "Service.Signature"))
		return ErrEmptyEmail
	}

	subject := "Welcome to our Newsletter!"
	body := "<h1>Thanks for subscribing!</h1><p>We are glad to have you with us.</p>"

	payload := mail.Message{
		To:      email,
		Subject: subject,
		Body:    body,
	}

	if err := se.mail.Process(payload); err != nil {
		slog.Error("mail delivery failed",
			slog.String("email", email),
			slog.Any("error", err),
		)
		return fmt.Errorf("%w: %v", ErrFailedToSend, err)
	}

	slog.Info("subscription email delivered successfully", slog.String("email", email))
	return nil
}
