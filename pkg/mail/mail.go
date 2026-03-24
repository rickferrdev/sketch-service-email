package mail

import (
	"log/slog"
	"strconv"

	"github.com/rickferrdev/sketch-service-email/config/env"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	dialer *gomail.Dialer
	jobs   chan Message
}

type Message struct {
	To      string
	Subject string
	Body    string
}

func New(env *env.Environment) (*Mail, error) {
	jobs := make(chan Message, 100)

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

	m := Mail{dialer: dialer, jobs: jobs}

	go m.worker()

	return &m, nil
}

func (ma *Mail) worker() {
	for message := range ma.jobs {
		slog.Debug("processing email sent", slog.String("to", message.To))
		if err := ma.Send(message); err != nil {
			slog.Error("email processing failed", slog.Any("error", err), slog.String("to", message.To))
		}
		slog.Info("email processed successfully in the background", slog.String("to", message.To))
	}
}

func (ma *Mail) Process(payload Message) error {
	ma.jobs <- payload
	slog.Debug("message added to queue", slog.String("to", payload.To))
	return nil
}

func (ma *Mail) Send(payload Message) error {
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
