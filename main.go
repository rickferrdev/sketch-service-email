package main

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/rickferrdev/sketch-service-email/config"
	"github.com/rickferrdev/sketch-service-email/internal/handlers"
	"github.com/rickferrdev/sketch-service-email/internal/services"
	"github.com/rickferrdev/sketch-service-email/pkg"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		pkg.Module,
		services.Module,
		handlers.Module,
		fx.Provide(
			NewApp,
			NewRouter,
		),
		fx.Invoke(Start),
	).Run()
}

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Sketch Email Service",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{"error": "internal error"})
		},
	})

	app.Use(logger.New())

	return app
}

func NewRouter(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
}

func Start(life fx.Lifecycle, app *fiber.App) {
	life.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info("starting fiber server", slog.String("port", "8080"))
			go func() {
				if err := app.Listen("0.0.0.0:8080"); err != nil {
					slog.Error("failed to start fiber server", slog.Any("error", err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("shutting down fiber server")
			return app.ShutdownWithContext(ctx)
		},
	})
}
