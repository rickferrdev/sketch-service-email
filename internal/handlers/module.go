package handlers

import (
	handlerSubscription "github.com/rickferrdev/sketch-service-email/internal/handlers/subscription"
	serviceSubscription "github.com/rickferrdev/sketch-service-email/internal/services/subscription"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers", fx.Provide(
	fx.Annotate(
		serviceSubscription.New,
		fx.As(new(handlerSubscription.Service)),
	),
), fx.Invoke(handlerSubscription.New))
