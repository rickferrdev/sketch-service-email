package mail

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/rickferrdev/sketch-service-email/config/env"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	dialer *gomail.Dialer
}

type Message struct {
	To      string
	Subject string
	Body    string
}

func New(env *env.Environment) (*Mail, error) {
	port, err := strconv.Atoi(env.MAIL_PORT)
	if err != nil {
		slog.Error("failed to parse mail port",
			slog.String("port", env.MAIL_PORT),
			slog.Any("error", err),
		)
		return nil, err
	}

	dialer := gomail.NewDialer(env.MAIL_HOST, port, env.MAIL_USER, env.MAIL_PASS)

	closer, err := dialer.Dial()
	if err != nil {
		slog.Error("failed to establish SMTP connection",
			slog.String("host", env.MAIL_HOST),
			slog.Int("port", port),
			slog.Any("error", err),
		)
		return nil, err
	}
	defer closer.Close()

	slog.Info("SMTP connection established successfully", slog.String("host", env.MAIL_HOST))

	return &Mail{dialer: dialer}, nil
}

func (ma *Mail) Send(ctx context.Context, payload Message) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "sketch-service-email@demomailtrap.co")
	mail.SetHeader("To", payload.To)
	mail.SetHeader("Subject", payload.Subject)
	mail.SetBody("text/html", payload.Body)

	if err := ma.dialer.DialAndSend(mail); err != nil {
		slog.Error("failed to send email",
			slog.String("to", payload.To),
			slog.String("subject", payload.Subject),
			slog.Any("error", err),
		)
		return err
	}

	slog.Debug("email sent successfully", slog.String("to", payload.To))
	return nil
}
