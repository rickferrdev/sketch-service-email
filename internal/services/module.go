package services

import (
	"github.com/rickferrdev/sketch-service-email/internal/services/subscription"
	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	subscription.New,
))
