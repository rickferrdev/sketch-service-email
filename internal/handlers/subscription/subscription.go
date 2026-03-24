package subscription

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

type Service interface {
	Signature(ctx context.Context, email string) error
}

func New(router fiber.Router, service Service) *Handler {
	handler := &Handler{
		service: service,
	}

	group := router.Group("/subs")
	group.Post("/", handler.Signature)

	return handler
}

type RequestSignatureDTO struct {
	Email string `json:"email"`
}

func (ha *Handler) Signature(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
	defer cancel()

	var body RequestSignatureDTO
	if err := c.Bind().Body(&body); err != nil {
		slog.Error("failed to bind request body", slog.Any("error", err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if body.Email == "" {
		slog.Warn("subscription attempt with empty email")
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "email is required",
		})
	}

	err := ha.service.Signature(ctx, body.Email)
	if err != nil {
		slog.Error("service failed to process signature",
			slog.String("email", body.Email),
			slog.Any("error", err),
		)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not process subscription",
		})
	}

	slog.Info("subscription successful", slog.String("email", body.Email))
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "subscription created successfully",
	})
}
