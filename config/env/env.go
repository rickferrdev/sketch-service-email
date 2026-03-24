package env

import (
	"log/slog"

	"github.com/rickferrdev/dotenv"
)

type Environment struct {
	MAIL_PORT string `env:"MAIL_PORT"`
	MAIL_HOST string `env:"MAIL_HOST"`
	MAIL_USER string `env:"MAIL_USERNAME"`
	MAIL_PASS string `env:"MAIL_PASSWORD"`
}

func New() (*Environment, error) {
	dotenv.Collect()

	var env Environment

	if err := dotenv.Unmarshal(&env); err != nil {
		slog.Error("failed to unmarshal environment variables", slog.Any("error", err))
		return nil, err
	}

	if env.MAIL_HOST == "" || env.MAIL_PORT == "" {
		slog.Error("missing critical mail configuration",
			slog.String("host", env.MAIL_HOST),
			slog.String("port", env.MAIL_PORT),
		)
	}

	slog.Info("environment variables loaded successfully")
	return &env, nil
}
